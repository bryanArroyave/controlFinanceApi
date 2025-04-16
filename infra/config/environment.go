package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables(env string) {

	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "LOCAL")
	}

	switch os.Getenv("ENV") {
	case "LOCAL":
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	default:
		return
	}
}
