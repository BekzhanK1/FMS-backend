#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Wait for the database to be ready
echo "Waiting for PostgreSQL to be ready at $PUBLIC_HOST:$DB_PORT..."
while ! nc -z $PUBLIC_HOST $DB_PORT; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 5
done

echo "PostgreSQL is up - running migrations..."

# Run migrations
go run cmd/migrate/main.go up

echo "Migrations completed successfully."
