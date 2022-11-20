# CRM

Crm backend service written in Golang.

## Run

Build Docker image: `docker build -t crm-backend . --build-arg GIT_TOKEN=<token> --build-arg GIT_USER=<git_user>`

Run Docker container: `docker run --network="<network>" --env-file <env-file> crm-backend`

Set environment variables (or put them iniline before `go run` command):
```bash
export LOG_LEVEL=debug
```

```bash
go run cmd/crm/main.go
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
```

## Requirements

[Application Requirements](https://review.udacity.com/#!/rubrics/4856/view)