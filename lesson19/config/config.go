package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func Load() Config {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("no env file found", err)
	}

	config := Config{}

	config.PostgresPort = cast.ToInt(getORReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresHost = cast.ToString(getORReturnDefaultValue("POSTGRES_HOST", ""))
	config.PostgresUser = cast.ToString(getORReturnDefaultValue("POSTGRES_USER", ""))
	config.PostgresPassword = cast.ToString(getORReturnDefaultValue("POSTGRES_PASSWORD", ""))
	config.PostgresDatabase = cast.ToString(getORReturnDefaultValue("POSTGRES_DB", ""))
	config.PostgresDatabase = cast.ToString(getORReturnDefaultValue("PORT", ""))

	return config
}

func getORReturnDefaultValue(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
