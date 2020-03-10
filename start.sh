#! /bin/bash

gorm-goose -path /mooncascade-task/db/ up

swag init

CompileDaemon --build="go build main.go" --command=./main