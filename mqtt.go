package main

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

func setupMQTTClient(config MQTTConfig) (MQTTClient, error) {
	mqttClient := MQTTClient{Client: nil, StateTopic: config.StateTopic}
	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcp://%s:%s", config.Host, config.Port))
	options.SetClientID("ecometer2mqtt")
	options.SetUsername(config.User)
	options.SetPassword(config.Password)

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return mqttClient, token.Error()
	}

	mqttClient.Client = client
	return mqttClient, nil
}

func (mqtt MQTTClient) setupHADiscovery() {
	HADevice := HADevice{
		Identifiers:  []string{"ecometer2mqtt"},
		Name:         "Ecometer2MQTT",
		Manufacturer: "Proteus",
		Model:        "Ecometer S",
		SwVersion:    "1.0.0",
	}

	temperature := HATopic{
		Field:          "temperature",
		Name:           "Temperatur",
		Icon:           "mdi:thermometer",
		SetRound:       true,
		Unit:           "°C",
		DeviceClass:    "temperature",
		StateClass:     "measurement",
		EntityCategory: "sensor",
	}

	distance := HATopic{
		Field:          "distance",
		Name:           "Distance",
		Icon:           "mdi:ruler",
		SetRound:       true,
		Unit:           "cm",
		DeviceClass:    "distance",
		EntityCategory: "sensor",
	}

	height := HATopic{
		Field:          "height",
		Name:           "Height",
		Icon:           "mdi:ruler",
		SetRound:       true,
		Unit:           "cm",
		DeviceClass:    "distance",
		EntityCategory: "sensor",
	}

	level := HATopic{
		Field:          "level",
		Name:           "Level",
		Icon:           "mdi:car-coolant-level",
		Unit:           "L",
		DeviceClass:    "volume_storage",
		EntityCategory: "sensor",
	}

	capacity := HATopic{
		Field:          "capacity",
		Name:           "Capacity",
		Icon:           "mdi:car-coolant-level",
		Unit:           "L",
		DeviceClass:    "volume_storage",
		EntityCategory: "sensor",
	}

	percent := HATopic{
		Field:          "percent",
		Name:           "Percent",
		Icon:           "mdi:car-coolant-level",
		SetRound:       true,
		Unit:           "%",
		EntityCategory: "sensor",
	}

	mqtt.sendHADiscovery(temperature, HADevice)
	mqtt.sendHADiscovery(distance, HADevice)
	mqtt.sendHADiscovery(height, HADevice)
	mqtt.sendHADiscovery(level, HADevice)
	mqtt.sendHADiscovery(capacity, HADevice)
	mqtt.sendHADiscovery(percent, HADevice)
}

func (mqtt MQTTClient) sendHADiscovery(haTopic HATopic, haDevice HADevice) {
	topic := fmt.Sprintf("homeassistant/%s/%s/%s/config", haTopic.EntityCategory, haTopic.Field, haDevice.Name)

	sensor := HASensor{
		Device:        haDevice,
		UniqueID:      haDevice.Identifiers[0] + "-" + haTopic.Field,
		ObjectID:      haDevice.Identifiers[0] + "-" + haTopic.Field,
		Name:          haTopic.Name,
		Icon:          haTopic.Icon,
		Unit:          haTopic.Unit,
		ForceUpdate:   true,
		StateTopic:    mqtt.StateTopic,
		ValueTemplate: fmt.Sprintf("{{ value_json.%s }}", haTopic.Field),
		DeviceClass:   haTopic.DeviceClass,
		StateClass:    haTopic.StateClass,
	}

	if haTopic.SetRound {
		sensor.ValueTemplate = fmt.Sprintf("{{ value_json.%s | round(2) }}", haTopic.Field)
	}

	sensorJson, err := json.Marshal(sensor)
	if err != nil {
		logrus.Error(err)
	}

	token := mqtt.Client.Publish(topic, 0, true, sensorJson)
	token.Wait()
}

func (mqtt MQTTClient) sendData(data EcometerData) {
	messageJSON, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}
	token := mqtt.Client.Publish(mqtt.StateTopic, 0, true, messageJSON)
	token.Wait()
}
