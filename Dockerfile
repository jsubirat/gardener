FROM balenalib/raspberrypi3-alpine-golang:latest

RUN install_packages libc-dev

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["gardener"]