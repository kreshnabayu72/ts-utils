package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ConfigOptions holds environment variables data as struct for convenient auto-complete.
//
// All variables will be loaded as string, it is up to you to parse it to any type you need.
type ConfigOptions struct {
	// SERVICE_NAME is the service's name. Defaults to "service"
	SERVICE_NAME string
	// ENV possible values: "production", "development", "local". Defaults to "local"
	ENV string
	// PORT is the service's exposed port. Defaults to "8080"
	PORT    string
	JWT_KEY string

	POSTGRES_USERNAME string
	POSTGRES_PASSWORD string

	// SSH_* is currently used exclusively for postgres connection.
	SSH_HOST     string
	SSH_USER     string
	SSH_PASSWORD string

	// Mailer specific configurations.
	MAILER_USER     string
	MAILER_PASSWORD string
	MAILER_HOST     string
	MAILER_PORT     string

	// Temporal configurations.
	// TEMPORAL_CLIENT_HOST is the temporal server host. Defaults to "localhost"
	TEMPORAL_CLIENT_HOST string
	// TEMPORAL_CLIENT_PORT is the temporal server port. Defaults to "7233"
	TEMPORAL_CLIENT_PORT string
}

// LoadConfig loads environment variables and returns as a struct.
func LoadConfig(filenames ...string) ConfigOptions {
	config := ConfigOptions{}

	// In containerized environments (Docker, Kubernetes), this function may return
	// an error (e.g., 'no such file or directory') because the .env file is not
	// included in the final image or not mounted at the expected path.
	err := godotenv.Load(filenames...)
	if err != nil {
		// Docker compose will inject env to os directly
		log.Print("Error while loading .env file\n", err)
	}

	if config.SERVICE_NAME = os.Getenv("SERVICE_NAME"); config.SERVICE_NAME == "" {
		config.SERVICE_NAME = "service"
	}

	if config.ENV = os.Getenv("ENV"); config.ENV == "" {
		config.ENV = "local"
	}

	if config.PORT = os.Getenv("PORT"); config.PORT == "" {
		config.PORT = "8080"
	}

	config.JWT_KEY = os.Getenv("JWT_KEY")

	config.POSTGRES_USERNAME = os.Getenv("POSTGRES_USERNAME")
	config.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")

	config.SSH_HOST = os.Getenv("SSH_HOST")
	config.SSH_USER = os.Getenv("SSH_USER")
	config.SSH_PASSWORD = os.Getenv("SSH_PASSWORD")

	config.MAILER_USER = os.Getenv("MAILER_USER")
	config.MAILER_PASSWORD = os.Getenv("MAILER_PASSWORD")
	config.MAILER_HOST = os.Getenv("MAILER_HOST")
	config.MAILER_PORT = os.Getenv("MAILER_PORT")

	if config.TEMPORAL_CLIENT_HOST = os.Getenv("TEMPORAL_CLIENT_HOST"); config.TEMPORAL_CLIENT_HOST == "" {
		config.TEMPORAL_CLIENT_HOST = "localhost"
	}

	if config.TEMPORAL_CLIENT_PORT = os.Getenv("TEMPORAL_CLIENT_PORT"); config.TEMPORAL_CLIENT_PORT == "" {
		config.TEMPORAL_CLIENT_PORT = "7233"
	}

	return config
}
