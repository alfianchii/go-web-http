FROM golang:1.23.3-alpine

WORKDIR /app

COPY . .

COPY ./docker/entrypoints/go.sh /app/docker/entrypoints/go.sh

ENTRYPOINT ["sh", "/app/docker/entrypoints/go.sh"]