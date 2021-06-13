package bootstrap

import (
	"fmt"
	"hotel-reservation/pkg/database"
)

// Run run application
func Run() error {
	fmt.Println("application started ...")
	err := database.Migrate()

	if err != nil {
		return err
	}

	return nil
}
