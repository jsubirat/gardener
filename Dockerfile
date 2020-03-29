FROM balenalib/raspberry-pi-alpine-python:3-latest

RUN install_packages gcc python3-dev musl-dev

RUN pip3 install RPI.GPIO adafruit-blinka adafruit-circuitpython-mcp3xxx

WORKDIR /app
COPY voltage.py .

CMD ["watch", "python3", "voltage.py"]


# FROM balenalib/raspberrypi3-alpine-golang:latest

# RUN install_packages libc-dev

# WORKDIR /go/src/app
# COPY . .

# RUN go get -d -v ./...
# RUN go install -v ./...

# CMD ["gardener"]