package env

import (
	"github.com/joho/godotenv"
	"os"
)

func GetFromDotENV(key string) string {
	err := godotenv.Load()
	if err != nil {
		return ""
	}
	return os.Getenv(key)
}

func GetFromOsENV(key string) string {
	return os.Getenv(key)
}
