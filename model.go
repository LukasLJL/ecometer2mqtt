package main

import mqtt "github.com/eclipse/paho.mqtt.golang"

type EcometerConfig struct {
	Height  uint16 `mapstructure:"height"`
	Offset  uint16 `mapstructure:"offset"`
	USBPort string `mapstructure:"usb_port"`
	Baud    int    `mapstructure:"baud"`
}

type MQTTConfig struct {
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	StateTopic string `mapstructure:"stateTopic"`
}

type Config struct {
	EcometerConfig EcometerConfig `mapstructure:"ecometer"`
	MQTTConfig     MQTTConfig     `mapstructure:"mqtt"`
}

type HATopic struct {
	Field          string `json:"field"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	Unit           string `json:"unit"`
	SetRound       bool   `json:"setRound"`
	DeviceClass    string `json:"device_class"`
	StateClass     string `json:"state_class"`
	EntityCategory string `json:"entity_category"`
}

type HADevice struct {
	Identifiers  []string `json:"identifiers"`
	Name         string   `json:"name"`
	Manufacturer string   `json:"manufacturer"`
	Model        string   `json:"model"`
	SwVersion    string   `json:"sw_version"`
}

type HASensor struct {
	Device        HADevice `json:"device"`
	UniqueID      string   `json:"unique_id"`
	ObjectID      string   `json:"object_id"`
	Name          string   `json:"name"`
	Icon          string   `json:"icon"`
	Unit          string   `json:"unit_of_measurement"`
	ForceUpdate   bool     `json:"force_update"`
	StateTopic    string   `json:"state_topic"`
	ValueTemplate string   `json:"value_template"`
	DeviceClass   string   `json:"device_class,omitempty"`
	StateClass    string   `json:"state_class,omitempty"`
}

type MQTTClient struct {
	Client     mqtt.Client
	StateTopic string
}

type EcometerData struct {
	Temperature float64 `json:"temperature"`
	Distance    uint16  `json:"distance"`
	Height      uint16  `json:"height"`
	Level       uint16  `json:"level"`
	Capacity    uint16  `json:"capacity"`
	Percent     float64 `json:"percent"`
}
