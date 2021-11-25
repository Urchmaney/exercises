FROM golang:latest

WORKDIR /usr/local/go/src/app

COPY . .


CMD ["tail", "-f", "/dev/null"]