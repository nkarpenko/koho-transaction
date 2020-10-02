# KOHO Transaction Tool
KOHO transaction tool technical assessment.

## Table of Contents
- [KOHO Transaction Tool](#koho-transaction-tool)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Go](#go)
  - [Usage](#usage)
    - [Commands](#commands)
    - [How to run with local config file](#how-to-run-with-local-config-file)
  - [Testing](#testing)
    - [Unit and Integration Tests](#unit-and-integration-tests)
    - [Benchmark Tests](#benchmark-tests)
- [Notes & TODO](#notes-and-todo)

## Installation
### Go
* Get all dependencies ```go get ./...```
* Run ```go run main.go```

## Usage
### Commands
* Review the config file in the root folder named **config.yml**. Defaults are already set for you. 
* To run the application, do ```go run main.go```
* **Run with custom config file**: run ```go run main.go -c filename.yml``` e.g. ```go run main.go -c config.local.yml```
* **Version** Get the tool version by running ```go run main.go version``` e.g.
```shell
$ go run main.go version
Koho user transaction tool v0.1
```
* **Limits** Get the tool limits by running ```go run main.go limits```. e.g.
```shell
$ go run main.go limits
User transaction limits:
        Max of $5000 can be loaded per day.
        Max of $20000 can be loaded per week.
        Max of 3 loads per day.
```
* **CLI Help** Get list of available commands and flags by running ```go run main.go help```. e.g.
```shell
go run main.go help     
Koho transaction validation tool.

Usage:
  koho-transaction [flags]
  koho-transaction [command]

Available Commands:
  help        Help about any command
  limits      Display user transaction limits.
  version     Display app version.

Flags:
  -c, --config string   Specify local configuration file. (default "config.yml")
  -h, --help            help for koho-transaction

Use "koho-transaction [command] --help" for more information about a command.
```

### How to run with local config file
* Create a local config file for example **config.local.yaml** in the root directory.
* This file is GIT ignored
* Run ```go run main.go -c config.local.yaml```

## Testing

### Unit and Integration Tests
To run all unit and integration tests, run:
``` shell
$ go test ./...
```

### Benchmark Tests
To run all benchmarks, run:
``` shell
$ go test -bench=.
```

# Notes and Todo
In a realistic production environment, this application would;
* Most likely be a HTTP REST API I would assume. It would leverage mux as an http server/router.
* rather than using a local cache variable for in memory storage, I would use a caching engine such as Redis or a nosql/sql solution to decouple data and memory storage from the application. This would also make it easier on searches vs sourting dates inside the application to enforce limits.
* Leverage a multi-worker based model with a queue system in place such as RabbitMQ or SQS to handle the sequence and integrity of transaction requests. The workers would poll for incoming messages, communicate with each other via channels and process the requests via go routines, being able to handle significantly more requests.
* Leverage docker/kube for local dev and deployments. Queues, cache and db would be stand alone services while the core application and workers would be deployed into containers.
