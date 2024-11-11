#!/bin/sh
set -e # make the script exit immediately if any command in the script returns a non-zero exit status

echo "run db migration"
/app/migrate -path /app/migration -database "${DB_SOURCE}" -verbose up

echo "start the app"
exec "$@" # execute the command provided in CMD of the Dockerfile