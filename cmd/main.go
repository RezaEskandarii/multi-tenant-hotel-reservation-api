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
// @Param Authorization header string true "With the bearer started"
// @Security  JWT
// @in header
// @name Authorization
// @tokenUrl http://127.0.0.1:8000/auth/signin
func main() {

	logger := applogger.New(nil)

	if err := kernel.Run(); err != nil {
		logger.LogError(err)
		os.Exit(1)
	}
}
