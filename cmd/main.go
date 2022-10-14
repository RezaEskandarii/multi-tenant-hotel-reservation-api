package main

import (
	"fmt"
	_ "reservation-api/docs"
	"reservation-api/internal/utils"
)

// @title Reservation API
// @version 1.0
// @description Swagger documentation for reservation API .
// @license.name MIT
// @BasePath /api/v1
func main() {

	//logger := applogger.New(nil)
	//
	//if err := kernel.Run(); err != nil {
	//	logger.LogError(err)
	//	os.Exit(1)
	//}

	bb := utils.Encrypt("REZAESS")
	fmt.Println(string(bb))
	fmt.Println("---------")
	fmt.Println(utils.Decrypt(bb))
}
