package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT            string
	QueueName       string
	MailgunDomain   string
	MailgunAPIKey   string
	SparkPostAPIKey string
}

// Global variable to hold the loaded configuration
var AppConfig *Config

func LoadConfig() *Config {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the configuration with environment variables
	AppConfig = &Config{
		PORT:            os.Getenv("PORT"),
		QueueName:       os.Getenv("QUEUE_NAME"),
		MailgunDomain:   os.Getenv("MAILGUN_DOMAIN"),
		MailgunAPIKey:   os.Getenv("MAILGUN_API_KEY"),
		SparkPostAPIKey: os.Getenv("SPARKPOST_API_KEY"),
	}

	return AppConfig
}
