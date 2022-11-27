package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mattiadevivo/crm/database"
	"github.com/mattiadevivo/crm/routes"
	"github.com/mattiadevivo/crm/variables"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// Setup logger depending on level string
func setupLogging(level string) {
	var logLevel zerolog.Level
	switch l := strings.ToLower(level); l {
	case "debug":
		logLevel = zerolog.DebugLevel
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	case "error":
		logLevel = zerolog.ErrorLevel
	default:
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)
}

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	setupLogging(os.Getenv(variables.LogLevel))
	// Server port
	port, isSet := os.LookupEnv(variables.ServerPort)
	if !isSet {
		port = "3000"
	}
	// Mock db migration
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_ADDRESS"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DBNAME"))
	if err := database.Connect(dsn); err != nil {
		log.Error().Err(err)
		panic(fmt.Sprint("Could not connect to the database", err))
	}
	if err := database.Migrate(); err != nil {
		log.Error().Err(err)
		panic(fmt.Sprint("Error during db migration: ", err))
	}
	// Setup api server
	app := fiber.New()
	routes.Setup(app)
	app.Listen(fmt.Sprintf(":%s", port))
	log.Info().Msg(fmt.Sprintf("Server started at port %s", port))
}
