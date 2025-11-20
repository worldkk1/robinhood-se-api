#!/bin/sh
set -e

goose -dir ./internal/database/migrations up

./api-server
