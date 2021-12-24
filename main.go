package main

import (
	"os"
	"reservation-api/internal/kernel"
	"reservation-api/pkg/applogger"
)

func main() {

	var port int

	logger := applogger.New(nil)
	kernel.SetCommands()
	err := kernel.Run(port)
	if err != nil {
		logger.LogInfo("exit ...")
		os.Exit(1)
	}
}
