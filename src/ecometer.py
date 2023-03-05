import datetime
import logging
import os
import struct
import threading

import serial

from mqtt import myMQTT

ECOMETER_SERIAL_PORT = os.environ['SERIAL_PORT']
HEIGHT = os.environ['TANK_HEIGHT']
OFFSET = os.environ['TANK_OFFSET']


class Ecometer(threading.Thread):
    ecometer_result = []

    def __init__(self):
        threading.Thread.__init__(self)

    def run(self):
        with serial.Serial(ECOMETER_SERIAL_PORT, 115200) as connection:
            while True:
                connection.reset_input_buffer()
                print("Waiting for data")
                data = connection.read(22)
                print("Data was receive")
                (header, _length, _command, _flags,
                    hour, minute, second, _start, _end,
                    temperature, distance, usable_level, capacity, _crc
                 ) = struct.unpack(">2shbb3bhhb4h", data)
                if header == b'SI':
                    payload = {
                        "time": f"{hour:02d}:{minute:02d}:{second:02d}",
                        "temperature": (temperature - 40 - 32) / 1.8,
                        "distance": distance,
                        "height":  int(HEIGHT) - distance + int(OFFSET),
                        "level": usable_level,
                        "capacity": capacity,
                        "percent": usable_level / capacity * 100.01,
                        "timestamp": datetime.datetime.now().timestamp()
                    }
                    print(payload)
                    mqtt = myMQTT()
                    mqtt.pushData(payload)
                    print("Data sent via MQTT")
