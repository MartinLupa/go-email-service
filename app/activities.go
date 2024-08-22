package app

import (
	"context"
	"log"

	"github.com/MartinLupa/go-email-service/config"
	sp "github.com/SparkPost/gosparkpost"
	"github.com/mailgun/mailgun-go/v4"
)

// Activities definition

func SendEmailViaMailgun(config *config.Config, from, subject, body, to string) (string, error) {
	// Docs: https://github.com/mailgun/mailgun-go
	mg := mailgun.NewMailgun(config.MailgunDomain, config.MailgunAPIKey)
	m := mg.NewMessage(
		from,
		subject,
		body,
		to,
	)
	id, _, err := mg.Send(context.Background(), m)

	return id, err
}

type SparkPostConfig struct {
	apiKey string
}

func SendEmailViaSparkPost(config *config.Config, from, subject, body, to string) (string, error) {
	// Docs: https://github.com/SparkPost/gosparkpost
	cfg := &sp.Config{
		BaseUrl:    "https://api.sparkpost.com",
		ApiKey:     config.SparkPostAPIKey,
		ApiVersion: 1,
	}
	var client sp.Client
	err := client.Init(cfg)
	if err != nil {
		log.Fatalf("SparkPost client init failed: %s\n", err)
	}

	tx := &sp.Transmission{
		Recipients: []string{to},
		Content: sp.Content{
			HTML:    "<p>" + body + "</p>",
			From:    from,
			Subject: subject,
		},
	}
	id, _, err := client.Send(tx)

	return id, err
}
