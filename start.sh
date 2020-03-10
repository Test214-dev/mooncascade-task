#! /bin/bash

echo "Checking database at ${POSTGRES_HOST} status..."

./wait-for $POSTGRES_HOST:5432 -t 50 -- echo "PostgreSQL is up, continuing..."

gorm-goose -path /mooncascade-task/db/ up

swag init

CompileDaemon --build="go build main.go" --command=./main