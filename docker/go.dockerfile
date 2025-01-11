FROM golang:1.23.3-alpine

RUN apk add --no-cache git && go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod .
RUN go mod tidy && go mod verify

COPY . .

COPY ./docker/entrypoints/go.sh /app/docker/entrypoints/go.sh

ENTRYPOINT ["sh", "/app/docker/entrypoints/go.sh"]