#!/usr/bin/env python3

from src.ecometer import Ecometer
from src.mqtt import myMQTT


def main():
    print("Hello World")

    my_ecometer = Ecometer()
    my_ecometer.start()


if __name__ == "__main__":
    main()
