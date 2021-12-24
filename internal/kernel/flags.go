package kernel

import (
	"flag"
	"fmt"
	"os"
	"reservation-api/internal/registery"
	"reservation-api/pkg/database"
)

// set application commands.
func loadFlags() {
	var migrate = false
	var seed = false
	flag.BoolVar(&migrate, "migrate", false, "migrate structs and struct changes to database.")
	flag.BoolVar(&seed, "seed", false, "seed default data from json file into database.")

	flag.Parse()

	if migrate {
		fmt.Println("migration started...")
		if err := database.Migrate(db); err != nil {
			logger.LogError(err)
			os.Exit(1)
		}
	}

	if seed {
		fmt.Println("seed started...")
		if err := registery.ApplySeed(db); err != nil {
			logger.LogError(err)
			os.Exit(1)
		}
	}
}
