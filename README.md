# KOHO Transaction Tool
KOHO transaction tool technical assessment. This simple application takes user transaction JSON input data, validates it based on configuration (derived from domain logic limitations), and processes those transactions to specify if they are valid or not. This tool leverages cobra and viper to make the executable application behave as a CLI 

## Table of Contents
- [KOHO Transaction Tool](#koho-transaction-tool)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Go](#go)
  - [Usage](#usage)
    - [Commands](#commands)
    - [How to run with local config file](#how-to-run-with-local-config-file)
  - [Testing](#testing)
    - [Unit Tests](#unit-tests)
    - [Integration Tests](#integration-tests)
    - [Benchmark Tests](#benchmark-tests)
  - [Documentation](#documentation)
  - [Debug](#debug)
- [Notes & TODO](#notes-and-todo)

## Installation
### Go
* Get all dependencies ```go get ./...```
* Run ```go run main.go``` to execute application.

## Usage
### Commands
* Review the config file in the root folder named **config.yml**. Defaults are already set for you. 
* To run the application, simply run ```go run main.go```
* **Specify Config File (Optional)**: run ```go run main.go -c filename.yml``` 
```shell
$ go run main.go -c config.local.yml
```
* **Version** Get the tool version by running ```go run main.go version```
```shell
$ go run main.go version
Koho user transaction tool v0.1
```
* **Limits** Get the tool limits by running ```go run main.go limits```. Limits are pulled from the config file.
```shell
$ go run main.go limits
User transaction limits:
        Max of $5000 can be loaded per day.
        Max of $20000 can be loaded per week.
        Max of 3 loads per day.
```
* **CLI Help** Get list of available commands and flags by running ```go run main.go help```
```shell
$ go run main.go help     
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

* **Have script output go to a file instead of regular stdout**
```shell
$ go run main.go > output.txt
```

### How to run with local config file
* Create a local config file for example **config.local.yml** in the root directory.
* This file is GIT ignored
* Run ```go run main.go -c config.local.yml```

## Testing

### Unit Tests
To run all unit and integration tests, run:
``` shell
$ go test ./...
```

### Integration Tests
TODO

### Benchmark Tests
TODO

## Documentation
To generate a `godoc` from the code, first make sure you have the base go tools which include godocs.
```shell
go get -u golang.org/x/tools/...
```
Generate the godoc:
```shell
godoc -http=:6060
```

Finally, navigate to the generated doc URL below to view package/file/interface/method docs+comments. 
* http://localhost:6060/pkg/github.com/nkarpenko/koho-transaction/

## Debug
To view the reason/message as to why a transaction did not get accepted, enable the ```message``` var inside of ```common/model/model.go``` by changing it's json tag to ```json:"message,omitempty"``` from ```json:"-"```. 

# Notes and Todo
In a realistic production environment, this application would;
* Most likely be a HTTP REST API I would assume. It would leverage mux as an http server/router.
* Rather than using a local cache variable for in memory storage, I would use a caching engine such as Redis or a nosql/sql solution to decouple data and memory storage from the application. This would also make it easier on searches vs sorting dates inside the application to enforce limits.
* Leverage a multi-worker based model with a queue system in place such as RabbitMQ or SQS to handle the sequence and integrity of transaction requests. The workers would poll for incoming messages, communicate with each other via channels and process the requests via go routines, being able to handle significantly more requests.
* Leverage docker/kube for local dev and deployments. Queues, cache and db would be stand alone services while the core application and workers would be deployed into containers.
* Include integration tests to cover as many user scenarios as possible utilizing Go Convey or any other Go BDD library.
* Include benchmarks within tests.
