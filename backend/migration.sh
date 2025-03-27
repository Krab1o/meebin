#!/bin/bash
source .env

export MIGRATION_DSN="host=$DB_HOST port=$PG_PORT dbname=$PG_DB user=$PG_USER password=$PG_PASSWORD"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v