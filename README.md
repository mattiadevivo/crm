# CRM

Crm backend service written in Golang.

## Run

Build Docker image: `docker build -t crm-backend . --build-arg GIT_TOKEN=<token> --build-arg GIT_USER=<git_user>`

Run Docker container: `docker run --network="<network>" --env-file <env-file> crm-backend`

Set environment variables (or put them iniline before `go run` command):
```bash
export LOG_LEVEL=debug
```

Create updated version of Swagger:
```
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g main.go --output docs/crm
```

Run the server:
```bash
go run main.go
```

## Environment Variables

Available log levels:
- panic (`zerolog.PanicLevel`, 5)
- fatal (`zerolog.FatalLevel`, 4)
- error (`zerolog.ErrorLevel`, 3)
- warn (`zerolog.WarnLevel`, 2)
- info (`zerolog.InfoLevel`, 1)
- debug (`zerolog.DebugLevel`, 0)
- trace (`zerolog.TraceLevel`, -1)

The application supports `debug` or `error` values for `LOG_LEVEL` env var, any other values (or no value) will be recognized as `info` level.

List of variables:
```bash
LOG_LEVEL=debug|error
PORT=3000
MYSQL_USER
MYSQL_PASSWORD
MYSQL_ADDRESS
MYSQL_PORT
MYSQL_DBNAME
```

## Project details

Backend service of a CRM web application. The server will support all of the functionalities:
- Getting a list of all customers
- Getting data for a single customer
- Adding a customer
- Updating a customer's information
- Removing a customer

## Requirements

[Application Requirements](https://review.udacity.com/#!/rubrics/4856/view)