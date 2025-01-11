#!/bin/sh
set -e # Exit immediately if a command exits with a non-zero status.

envsubst '${APP_URL} ${APP_PORT}' < /etc/nginx/conf.d/default.conf.template > /etc/nginx/conf.d/default.conf
exec "$@"