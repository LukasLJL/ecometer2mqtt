import json
import os

import paho.mqtt.client as mqtt


class myMQTT():
    def discoverTime(self, device):
        topic = "homeassistant/sensor/ecometer2mqtt/time/config"
        payload = {
            "device": device,
            "name": "Time",
            "unique_id": "eco2mqtt_time",
            "object_id": "eco2mqtt_time",
            "unit_of_measurement": "",
            "device_class": "time",
            "state_topic": "ecometer2mqtt/state",
            "value_template": "{{ value_json.time }}",
        }
        self.client.publish(topic, json.dumps(payload))
        
    def discoverTemperature(self, device):
        topic = "homeassistant/sensor/ecometer2mqtt/temp/config"
        payload = {
            "device": device,
            "name": "Temperature",
            "unique_id": "eco2mqtt_temp",
            "object_id": "eco2mqtt_temp",
            "unit_of_measurement": "°C",
            "device_class": "temperature",
            "state_topic": "ecometer2mqtt/state",
            "value_template": "{{ value_json.temperature }}",
        }
        self.client.publish(topic, json.dumps(payload))
        
    def discoverDistance(self, device):
        topic = "homeassistant/sensor/ecometer2mqtt/distance/config"
        payload = {
            "device": device,
            "name": "Distance",
            "unique_id": "eco2mqtt_distance",
            "object_id": "eco2mqtt_distance",
            "unit_of_measurement": "cm",
            "state_topic": "ecometer2mqtt/state",
            "value_template": "{{ value_json.distance }}",
        }
        self.client.publish(topic, json.dumps(payload))
        
    def discoverHeight(self, device):
        topic = "homeassistant/sensor/ecometer2mqtt/height/config"
        payload = {
            "device": device,
            "name": "Height",
            "unique_id": "eco2mqtt_height",
            "object_id": "eco2mqtt_height",
            "unit_of_measurement": "cm",
            "state_topic": "ecometer2mqtt/state",
            "value_template": "{{ value_json.height }}",
        }
        self.client.publish(topic, json.dumps(payload))

    def discoverLevel(self, device):
        topic = "homeassistant/sensor/ecometer2mqtt/level/config"
        payload = {
            "device": device,
            "name": "Level",
            "unique_id": "eco2mqtt_level",
            "object_id": "eco2mqtt_level",
            "unit_of_measurement": "Liter",
            "state_topic": "ecometer2mqtt/state",
            "value_template": "{{ value_json.level }}",
        }
        self.client.publish(topic, json.dumps(payload))

    def discoverCapacity(self, device):
        topic = "homeassistant/sensor/ecometer2mqtt/capacity/config"
        payload = {
            "device": device,
            "name": "Capacity",
            "unique_id": "eco2mqtt_capacity",
            "object_id": "eco2mqtt_capacity",
            "unit_of_measurement": "Liter",
            "state_topic": "ecometer2mqtt/state",
            "value_template": "{{ value_json.capacity }}",
        }
        self.client.publish(topic, json.dumps(payload))

    def discoverPercent(self, device):
        topic = "homeassistant/sensor/ecometer2mqtt/percent/config"
        payload = {
            "device": device,
            "name": "Percent",
            "unique_id": "eco2mqtt_percent",
            "object_id": "eco2mqtt_percent",
            "unit_of_measurement": "%",
            "state_topic": "ecometer2mqtt/state",
            "value_template": "{{ value_json.percent }}",
        }
        self.client.publish(topic, json.dumps(payload))

    def pushData(self, payload):
        topic = "ecometer2mqtt/state"
        self.client.publish(topic, json.dumps(payload))

    def pushStart(self):
        topic = "ecometer2mqtt/info"
        self.client.publish(topic, json.dumps({
            "msg": "Starting ecometer2mqtt",
        }))

    def __init__(self):
        client = mqtt.Client("ecometer2mqtt-python")
        client.username_pw_set(os.environ['MQTT_USER'], os.environ['MQTT_PASSWORD'])
        client.connect(os.getenv('MQTT_BROKER'), int(os.environ['MQTT_PORT']))
        self.client = client

        device = {
            "identifiers": "ecometer2mqtt-python",
            "model": "Ecometer S",
            "manufacturer": "Proteus",
            "name": "Ecometer2MQTT",
            "sw_version": "0.0.2"
        }

        self.discoverTime(device)
        self.discoverTemperature(device)
        self.discoverDistance(device)
        self.discoverHeight(device)
        self.discoverLevel(device)
        self.discoverCapacity(device)
        self.discoverPercent(device)
