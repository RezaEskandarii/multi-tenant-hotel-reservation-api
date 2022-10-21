package kernel

import (
	"flag"
	"fmt"
	"os"
	"reservation-api/pkg/database"
)

// set application commands.
func loadFlags() {

	var setup = false
	var migrate = false

	flag.BoolVar(&setup, "setup", false, "create and migrate tenant databases")
	flag.BoolVar(&migrate, "migrate", false, "migrate database changes")
	flag.Parse()

	if setup {
		fmt.Println("setup started...")
		if err := database.SetUp(); err != nil {

			logger.LogError(err)
			os.Exit(1)
		}
	}

	if migrate {
		fmt.Println("migration started...")
		if err := database.Migrate(); err != nil {

			logger.LogError(err)
			os.Exit(1)
		}
	}

}
