# Ecometer2MQTT
This project was forked from [github.com/seranoo/ecometer](https://github.com/seranoo/ecometer), to provide the Home Assistant device discovery feature and support for docker.

## Getting started
Here is an example docker-compose file which you can use:
````yaml
version: "3.8"
services:
  ecometer2mqtt:
    image: ghcr.io/lukasljl/ecometer2mqtt:latest
    container_name: ecometer2mqtt
    restart: unless-stopped
    devices:
      - /dev/ttyUSB0
    env_file:
      - ./config.env
````
You will have to set your correct environment variables in a file and use the ``env_file`` feature or you can set every environment variable manually.

## Environment Variables
- ``SERIAL_PORT`` defines the serial port which is used to communicate with the ecometer device.
- ``TANK_HEIGHT`` (optional) provide the height of your tank to calculate the height of the water level.
- ``TANK_OFFSET`` (optional) provide the offset of your sensor to calculate the height of the water level.
- ``MQTT_BROKER`` MQTT host-ip to connect to your mqtt broker.
- ``MQTT_PORT`` MQTT port to connect to your mqtt broker, the default is 1883.
- ``MQTT_USER`` MQTT user to connect to your mqtt broker.
- ``MQTT_PASSWORD``MQTT password of your user to connect to your mqtt broker.