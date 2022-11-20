package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
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

func main() {
	setupLogging(os.Getenv(variables.LogLevel))

	port, isSet := os.LookupEnv(variables.ServerPort)
	if !isSet {
		port = "3000"
	}
	app := fiber.New()
	routes.Setup(app)

	app.Listen(fmt.Sprintf(":%s", port))
	log.Info().Msg(fmt.Sprintf("Server started at port %s", port))
}
