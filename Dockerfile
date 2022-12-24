FROM golang:1.18-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

LABEL maintainer="Reza Eskandari"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY ./data .
COPY ./resources .

RUN go build ./cmd/main.go

EXPOSE 8080

CMD ["./main","-migrate","-setup"]