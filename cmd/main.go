package main

import (
	"fmt"
	"os"
	"reservation-api/cmd/build"
	_ "reservation-api/docs"
	"reservation-api/internal/app"
	"reservation-api/pkg/applogger"
	"time"
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
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey Bearer
//
func main() {

	logger := applogger.New(nil)
	logger.LogInfo(fmt.Sprintf("Application started at: {%s}", time.Now().String()))
	defer func() {
		if r := recover(); r != nil {
			logger.LogInfoJSON(r)
		}
	}()

	BuildInfo.Print()
	if err := app.Run(); err != nil {
		logger.LogError(err)
		os.Exit(1)
	}

}
