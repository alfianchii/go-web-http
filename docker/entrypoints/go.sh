#!/bin/sh
set -e # Exit immediately if a command exits with a non-zero status.

# Run go mod tidy, go mod download, and go mod verify
go mod tidy && go mod download && go mod verify
air -c air.toml