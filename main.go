package main

import (
	"os"
	"reservation-api/internal/kernel"
	"reservation-api/pkg/applogger"
)

func main() {

	logger := applogger.New(nil)

	if err := kernel.Run(); err != nil {
		logger.LogError(err)
		os.Exit(1)
	}
}
