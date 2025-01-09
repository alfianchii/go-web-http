FROM golang:1.23.3-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY ./docker/entrypoints/go.sh /app/docker/entrypoints/go.sh

ENTRYPOINT ["sh", "/app/docker/entrypoints/go.sh"]
