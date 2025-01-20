#!/bin/sh
set -e # Exit immediately if a command exits with a non-zero status.

APP_ENV=${APP_ENV:-production}

echo "Running in ${APP_ENV} environment"

if [ "$APP_ENV" = "development" ]; then
  echo "Setting up for development..."

  # Air
  if ! command -v air >/dev/null 2>&1; then
    echo "Air not found, installing..."
    apk add --no-cache git && go install github.com/air-verse/air@latest
  else
    echo "Air already installed."
  fi

  # go.sum file
  if [ -f "go.sum" ]; then
    echo "'go.sum' exists, skipping go mod tidy and go mod download."
  else
    echo "'go.sum' not found, running go mod tidy and go mod download."
    go mod tidy && go mod download && go mod verify
  fi
  
  air -c air.toml
elif [ "$APP_ENV" = "production" ]; then
  echo "Setting up for production..."
  apk add --no-cache git gettext
  
  # go.sum file
  if [ -f "go.sum" ]; then
    echo "'go.sum' exists, skipping go mod tidy and go mod download."
  else
    echo "'go.sum' not found, running go mod tidy and go mod download."
    go mod tidy && go mod download && go mod verify
  fi

  # main file
  if [ -f "main" ]; then
    echo "main exists, skipping go build."
  else
    echo "'main' not found, running go build."
    go build -buildvcs=false -o main .
  fi

  ./main
fi