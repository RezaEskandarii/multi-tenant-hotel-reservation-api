package main

import (
	"os"
	"reservation-api/cmd/build"
	_ "reservation-api/docs"
	"reservation-api/internal/kernel"
	"reservation-api/pkg/applogger"
)

var (
	BuildInfo = build.Build{}
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
	defer func() {
		if r := recover(); r != nil {
			logger.LogInfoJSON(r)
		}
	}()

	BuildInfo.Print()

	if err := kernel.Run(); err != nil {
		logger.LogError(err)
		os.Exit(1)
	}
}
