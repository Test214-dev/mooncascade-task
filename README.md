# mooncascade-task
A project for fictional sport event

## How to build and start:
Run `docker-compose up -d` and wait. It might take a while for postgres container to set up.
App listens on port 8080

## Documentation:
* This file
* Visit swagger at http://localhost:8080/swagger/index.html

## Command-line tool:
* Go to cli directory
* Build it, `go build`
* Run `./cli --help` to see available commands
* If you want to deploy app on server other than localhost, new host address must be put into `cli/config` file
