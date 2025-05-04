#!/bin/bash

set -e

echo "Running migrations..."
make db-migrate-up

echo "Starting server..."
exec ./main.exe