#!/bin/sh
set -e # Exit immediately if a command exits with a non-zero status.

APP_ENV=${APP_ENV:-production}

echo "Running in ${APP_ENV} environment"

if [ "$APP_ENV" = "development" ]; then
  echo "Setting up for development..."
  apk add --no-cache git && go install github.com/air-verse/air@latest
  go mod tidy && go mod verify
  air -c air.toml
elif [ "$APP_ENV" = "production" ]; then
  echo "Setting up for production..."
  apk add --no-cache git gettext
  go mod tidy && go mod verify
  go build -buildvcs=false -o main .
  ./main
else
  echo "Unknown APP_ENV value: ${APP_ENV}"
  exit 1
fi