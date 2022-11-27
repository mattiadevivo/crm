# CRM

Crm backend service written in Golang (used as Udacity Golang project work) that uses Mysql as database for customers data usage together with **gorm** as ORM.
The web framework used in this project is **gofiber**, a really fast and simple one.
As logging library `zerlog` has been used.

**Project details**:
Backend service of a CRM web application. The server will support all of the functionalities:
- Getting a list of all customers
- Getting data for a single customer
- Adding a customer
- Updating a customer's information
- Removing a customer

When running the project the first time database schema will be automatically created and some dummy data stored.

At `http://localhost:<port>/` you'll find a static html file describing the project.
At `http://localhost:<port>/swagger/index.html` you'll find Swagger documentation.

## Run

Build Docker image: `docker build -t crm-backend . --build-arg GIT_TOKEN=<token> --build-arg GIT_USER=<git_user>`

Run Docker container: `docker run --network="<docker-compose-mysql-network>" --env-file <container.env/your-env-file> -p "127.0.0.1:<port-on-host>:<container-server-port>" crm-backend`

Create updated version of Swagger:
```
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g main.go --output docs/crm
```

Create Mysql Docker container:
```bash
docker-compose up -d
```

Run the server:
```bash
# export env vars
source test.env 
# download dependencies
go mod tidy
# execute backend server
go run main.go
```

Check `/postman` folder and load inside Postman the `collection.json` file to have some ready-to-use requests to run against the server.

## Mysql administration

Run the command to login in mysql cli inside Mysql Docker container:
```bash
docker exec -ti mysql mysql -u crm-backend -p
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

List of variables (see them under `/variables/env.go`):
```bash
LOG_LEVEL=debug|error
PORT=3000
MYSQL_USER=crm-backend
MYSQL_PASSWORD=crm-backend
MYSQL_ADDRESS=localhost
MYSQL_PORT=3306
MYSQL_DBNAME=crm
```

## Requirements

[Application Requirements](https://review.udacity.com/#!/rubrics/4856/view)

## References

[gofiber](https://gofiber.io/)
[gorm](https://gorm.io/)