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

	err := kernel.Run(port)
	if err != nil {
		applogger.LogInfo("exit ...")
		os.Exit(1)
	}

}
