package main

import (
	"flag"
	"hotel-reservation/internal/kernel"
	"hotel-reservation/pkg/applogger"
	"os"
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
