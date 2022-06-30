package kernel

import (
	"flag"
	"fmt"
	"os"
	"reservation-api/pkg/database"
)

// set application commands.
func loadFlags() {
	var migrate = false
	var seed = false
	flag.BoolVar(&migrate, "migrate", false, "migrate structs and struct changes to database.")
	flag.BoolVar(&seed, "seed", false, "seed default data from json file into database.")
	flag.Parse()

	var tx = db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if migrate {
		fmt.Println("migration started...")
		if err := database.Migrate(tx); err != nil {
			logger.LogError(err)
			os.Exit(1)
		}
	}

	if seed {
		fmt.Println("seed started...")
		if err := database.ApplySeed(tx); err != nil {
			logger.LogError(err)
			os.Exit(1)
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
	}
}
