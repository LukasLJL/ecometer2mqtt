FROM python:3.10-alpine

WORKDIR /ecometer2mqtt
COPY ./src/main.py .
COPY ./src/ecometer.py .
COPY ./src/mqtt.py .
COPY ./src/requirements.txt .

RUN pip3 install --no-cache-dir -r requirements.txt

ENTRYPOINT [ "python3", "main.py"]
