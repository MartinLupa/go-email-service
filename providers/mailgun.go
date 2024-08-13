package providers

import (
	"context"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunProvider struct {
	domain string
	apiKey string
}

func NewMailgunProvider(domain, apiKey string) *MailgunProvider {
	return &MailgunProvider{
		domain: domain,
		apiKey: apiKey,
	}
}

func (p *MailgunProvider) SendEmail(from, subject, body, to string) (string, error) {
	// Docs: https://github.com/mailgun/mailgun-go
	mg := mailgun.NewMailgun(p.domain, p.apiKey)
	m := mg.NewMessage(
		from,
		subject,
		body,
		to,
	)
	id, _, err := mg.Send(context.Background(), m)

	return id, err
}
