package tenant_connection_string_resolver

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func GetDbLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Disable color
		},
	)
}
