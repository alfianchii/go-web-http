#!/bin/sh
set -e # Exit immediately if a command exits with a non-zero status.

go mod tidy && go mod download && go mod verify
go build -o main .
./main