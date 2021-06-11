package main

import (
	"fmt"
	"hotel-reservation/config"
)

func main() {

	cfg, err := config.NewConfig()

	if err == nil {
		fmt.Println(cfg)
	}
}
