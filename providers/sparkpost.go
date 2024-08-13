package providers

import (
	"log"

	sp "github.com/SparkPost/gosparkpost"
)

type SparkPostProvider struct {
	apiKey string
}

func NewSparkPostProvider(apiKey string) *SparkPostProvider {
	return &SparkPostProvider{apiKey: apiKey}
}

func (p *SparkPostProvider) SendEmail(from, subject, body, to string) (string, error) {
	// Docs: https://github.com/SparkPost/gosparkpost
	cfg := &sp.Config{
		BaseUrl:    "https://api.sparkpost.com",
		ApiKey:     p.apiKey,
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
