FROM golang:1.23.3-alpine

RUN apk add --no-cache git gettext

WORKDIR /app

COPY . .

COPY ./docker/entrypoints/go.prod.sh /app/docker/entrypoints/go.prod.sh

ENTRYPOINT ["sh", "/app/docker/entrypoints/go.prod.sh"]