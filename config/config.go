package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ConfigOptions holds environment variables data as struct for convenient auto-complete.
type ConfigOptions struct {
	POSTGRES_USERNAME string
	POSTGRES_PASSWORD string

	// SSH_* is currently used exclusively for postgres connection.
	SSH_HOST     string
	SSH_USER     string
	SSH_PASSWORD string

	MAILER_USER     string
	MAILER_PASSWORD string
	MAILER_HOST     string
	MAILER_PORT     string
}

// LoadConfig loads environment variables and returns as a struct.
func LoadConfig(filenames ...string) ConfigOptions {
	config := ConfigOptions{}

	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatal("Error while loading .env file", err)
	}

	config.POSTGRES_USERNAME = os.Getenv("POSTGRES_USERNAME")
	config.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")

	config.SSH_HOST = os.Getenv("SSH_HOST")
	config.SSH_USER = os.Getenv("SSH_USER")
	config.SSH_PASSWORD = os.Getenv("SSH_PASSWORD")

	config.MAILER_USER = os.Getenv("MAILER_USER")
	config.MAILER_PASSWORD = os.Getenv("MAILER_PASSWORD")
	config.MAILER_HOST = os.Getenv("MAILER_HOST")
	config.MAILER_PORT = os.Getenv("MAILER_PORT")

	return config
}
