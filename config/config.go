package config

import (
	"os"
)

type Config struct {
	PORT            string
	MailgunDomain   string
	MailgunAPIKey   string
	SparkPostAPIKey string
}

func LoadConfig() *Config {
	return &Config{
		PORT:            os.Getenv("PORT"),
		MailgunDomain:   os.Getenv("MAILGUN_DOMAIN"),
		MailgunAPIKey:   os.Getenv("MAILGUN_API_KEY"),
		SparkPostAPIKey: os.Getenv("SPARKPOST_API_KEY"),
	}
}
