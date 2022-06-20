package main

import (
	"os"
	_ "reservation-api/docs"
	"reservation-api/internal/kernel"
	"reservation-api/pkg/applogger"
)

// @title Reservation API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /api/v1
func main() {

	logger := applogger.New(nil)

	if err := kernel.Run(); err != nil {
		logger.LogError(err)
		os.Exit(1)
	}
}
