#!/bin/sh
set -e # make the script exit immediately if any command in the script returns a non-zero exit status

echo "run db migration"
source /app/app.env # Overwrite DB_SOURCE to point to AWS RDS instance
/app/migrate -path /app/migration -database "${DB_SOURCE}" -verbose up

echo "start the app"
exec "$@" # execute the command provided in CMD of the Dockerfile or the command attribute in docker-compose.yaml