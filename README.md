# mooncascade-task
A project for fictional sport event. The goal is to track timing points for athletes passing either finish corridor point, or finish line point.

## Athlete asset
* ChipID - ID of athlete chip (UUID)
* FullName - athlete full name
* StartNumber - athlete start number

## Timing asset
* TimingID - ID of timing asset (UUID)
* PointID - ID of point that was passed, either 'finish line' or 'finish corridor'
* Timestamp - timestamp

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

CLI supports 6 commands overall: (as well as the entire backend API):
* Create an athlete
* Add a timing for an athlete
* List all athletes
* List all timings
* Get a timing by ID
* Get an athlete by ID
