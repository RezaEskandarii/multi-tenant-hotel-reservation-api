package main

import (
	"fmt"
	"hotel-reservation/internal/bootstrap"
	"os"
)

func main() {

	err := bootstrap.Run()

	if err != nil {
		fmt.Println("exit ...")
		os.Exit(1)
	}
}
