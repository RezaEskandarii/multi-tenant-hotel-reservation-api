package main

import (
	"os"
	_ "reservation-api/docs"
	"reservation-api/internal/kernel"
	"reservation-api/pkg/applogger"
)

// @title Reservation API
// @version 1.0
// @description Swagger documentation for reservation API .
// @license.name MIT
// @BasePath /api/v1
func main() {

	logger := applogger.New(nil)

	if err := kernel.Run(); err != nil {
		logger.LogError(err)
		os.Exit(1)
	}
}
