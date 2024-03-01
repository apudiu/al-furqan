package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	// Return the value of the variable
	return os.Getenv(key)
}

func EnvDebug() bool {
	return Config("DEBUG") == "1"
}
func EnvKey() string {
	return Config("KEY")
}
func EnvDbHost() string {
	return Config("DB_HOST")
}
func EnvDbPort() string {
	return Config("DB_PORT")
}
func EnvDbName() string {
	return Config("DB_NAME")
}
func EnvDbUser() string {
	return Config("DB_USER")
}
func EnvDbPass() string {
	return Config("DB_PASS")
}
