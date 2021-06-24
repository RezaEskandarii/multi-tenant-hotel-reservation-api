package main

import (
	"flag"
	"hotel-reservation/internal/bootstrap"
	"hotel-reservation/pkg/application_loger"
	"os"
)

func main() {

	var port int

	flag.IntVar(&port, "port", 8080, "application port")
	flag.Parse()

	err := bootstrap.Run(port)
	if err != nil {
		application_loger.LogInfo("exit ...")
		os.Exit(1)
	}
}
