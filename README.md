# Ecometer2MQTT
This project was forked from [github.com/seranoo/ecometer](https://github.com/seranoo/ecometer), to provide the Home Assistant device discovery feature and support for docker.

## Getting started
Here is an example docker-compose file which you can use:
````yaml
services:
  ecometer2mqtt:
    image: ghcr.io/lukasljl/ecometer2mqtt:latest
    container_name: ecometer2mqtt
    restart: unless-stopped
    devices:
      - /dev/ttyUSB0
    environment:
      - USB_PORT=/dev/ttyUSB
    volumes:
      - ./ecometer2mqtt/config.yaml:/app/config.yaml
````
You will have to set your correct configuration within the ``config.yaml``.

## Configuration Parameters
````yaml
ecometer:
  usb_port: /dev/ttyUSB0 # defines the serial port which is used to communicate with the ecometer device.
  baud: 115200 # sets the baud rate for the serial usb connection
  height: 145 # (optional) provide the height of your tank to calculate the height of the water level.
  offset: 21 # (optional) provide the offset of your sensor to calculate the height of the water level.
mqtt:
  host: myMqttBroker.local # MQTT host to connect to your mqtt broker.
  port: 1883 # MQTT port to connect to your mqtt broker, the default is 1883.
  user: mqtt-user # MQTT user to connect to your mqtt broker.
  password: mqtt-password # MQTT password of your user to connect to your mqtt broker.
  stateTopic: ecometer2mqtt/state # MQTT state topic where the results of the ecometer will be send.
````
