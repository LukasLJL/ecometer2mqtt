package main

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tarm/serial"
)

func main() {
	logrus.Info("starting ecometer2mqtt...")

	config, err := getConfig()
	if err != nil {
		logrus.Error("Error getting config: ", err)
		return
	}

	client, err := setupMQTTClient(config.MQTTConfig)
	if err != nil {
		logrus.Error("Error setting up MQTT client: ", err)
		return
	}

	client.setupHADiscovery()
	readEcometer(config.EcometerConfig, client)
}

func getConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return Config{}, err
	}

	var config Config
	err = viper.Unmarshal(&config)

	if err != nil {
		return Config{}, err
	}

	if val := viper.GetString("USB_PORT"); val != "" {
		config.EcometerConfig.USBPort = val
	}

	return config, nil
}

func readEcometer(cfg EcometerConfig, mqtt MQTTClient) {
	connection, err := serial.OpenPort(&serial.Config{Name: cfg.USBPort, Baud: cfg.Baud})
	if err != nil {
		logrus.Error("Error opening serial port:", err)
		return
	}
	defer connection.Close()

	for {
		logrus.Debug("Waiting for data")
		data := make([]byte, 22)
		_, err := connection.Read(data)
		if err != nil {
			logrus.Error("Error reading from serial port:", err)
			continue
		}
		logrus.Debug("Data was received")

		var header [2]byte
		var length uint16
		var command, flags byte
		var hour, minute, second uint8
		var start, end uint16
		var temperature byte
		var distance, usableLevel, capacity, crc uint16

		buf := bytes.NewReader(data)
		err = binary.Read(buf, binary.BigEndian, &header)
		err = binary.Read(buf, binary.BigEndian, &length)
		err = binary.Read(buf, binary.BigEndian, &command)
		err = binary.Read(buf, binary.BigEndian, &flags)
		err = binary.Read(buf, binary.BigEndian, &hour)
		err = binary.Read(buf, binary.BigEndian, &minute)
		err = binary.Read(buf, binary.BigEndian, &second)
		err = binary.Read(buf, binary.BigEndian, &start)
		err = binary.Read(buf, binary.BigEndian, &end)
		err = binary.Read(buf, binary.BigEndian, &temperature)
		err = binary.Read(buf, binary.BigEndian, &distance)
		err = binary.Read(buf, binary.BigEndian, &usableLevel)
		err = binary.Read(buf, binary.BigEndian, &capacity)
		err = binary.Read(buf, binary.BigEndian, &crc)
		if err != nil {
			logrus.Error("Error unpacking data:", err)
		}

		if logrus.GetLevel() == logrus.DebugLevel {
			logrus.Debug("Header:", string(header[:]))

			remainingBytes, _ := io.ReadAll(buf)
			logrus.Debug("Buf:", string(remainingBytes))

			logrus.Debug("Data:", data)
		}

		if string(header[:]) == "SI" {
			data := EcometerData{
				Temperature: (float64(temperature) - 40 - 32) / 1.8,
				Distance:    distance,
				Height:      cfg.Height - distance + cfg.Offset,
				Level:       usableLevel,
				Capacity:    capacity,
				Percent:     float64(usableLevel) / float64(capacity) * 100.01,
			}
			logrus.Debug(data)
			mqtt.sendData(data)
		}
	}
}
