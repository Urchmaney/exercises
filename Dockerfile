FROM golang:latest

WORKDIR /go/src/app

COPY . .


CMD ["tail", "-f", "/dev/null"]