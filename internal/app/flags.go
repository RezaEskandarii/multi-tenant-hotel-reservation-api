package app

import (
	"flag"
	"fmt"
	"os"
	"reservation-api/pkg/multi_tenancy_database"
)

// set application commands.
func loadFlags() {

	var setup = false
	var migrate = false

	flag.BoolVar(&setup, "setup", false, "create and migrate tenant databases")
	flag.BoolVar(&migrate, "migrate", false, "migrate multi_tenancy_database changes")
	flag.Parse()

	if setup {
		fmt.Println("setup started...")
		if err := multi_tenancy_database.ClientSetUp(); err != nil {

			logger.LogError(err)
			os.Exit(1)
		}
	}

	if migrate {
		fmt.Println("migration started...")
		if err := multi_tenancy_database.Migrate(); err != nil {

			logger.LogError(err)
			os.Exit(1)
		}
	}

}
