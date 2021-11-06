package main

import (
	"flag"
	"os"
	"reservation-api/internal/kernel"
	"reservation-api/pkg/applogger"
)

func main() {

	var port int

	flag.IntVar(&port, "port", 8080, "application port")
	flag.Parse()

	logger := applogger.New(nil)

	err := kernel.Run(port)
	if err != nil {
		logger.LogInfo("exit ...")
		os.Exit(1)
	}
}
