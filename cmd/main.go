package main

import (
	"os"
	"reservation-api/internal/kernel"
	"reservation-api/internal/models"
	"reservation-api/internal/services/common_services"
	"reservation-api/pkg/applogger"
)

func main() {

	logger := applogger.New(nil)

	r := common_services.NewReportService(nil)

	x := make([]models.User, 0)
	x = append(x, models.User{Username: "hhhh", FirstName: "eskandari"})

	r.ExportToExcel(x, "fa")

	if err := kernel.Run(); err != nil {
		logger.LogError(err)
		os.Exit(1)
	}
}
