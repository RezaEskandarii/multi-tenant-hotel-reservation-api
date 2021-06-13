package bootstrap

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hotel-reservation/internal/registery"
	"hotel-reservation/pkg/application_loger"
	"hotel-reservation/pkg/database"
)

// Run run application
func Run(port int) error {
	fmt.Println("application started ...")

	defer func() {
		if r := recover(); r != nil {
			application_loger.LogError(r)
			return
		}
	}()

	db, err := database.GetDb()

	if err != nil {
		return err
	}

	err = database.Migrate(db)

	if err != nil {
		return err
	}

	portStr := fmt.Sprintf(":%d", port)

	e := echo.New()
	router := e.Group("/v1")

	registery.RegisterServices(db, router)

	e.Logger.Fatal(e.Start(portStr))

	return nil
}
