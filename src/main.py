#!/usr/bin/env python3

from ecometer import Ecometer
from mqtt import myMQTT


def main():
    print("Hello World")

    my_ecometer = Ecometer()
    my_ecometer.start()


if __name__ == "__main__":
    main()
