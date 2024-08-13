package service

import (
	"log"
	"time"

	"github.com/MartinLupa/go-email-service/providers"
)

type EmailService struct {
	providers []providers.EmailProvider
	retries   int
}

func NewEmailService(providers []providers.EmailProvider, retries int) *EmailService {
	return &EmailService{
		providers: providers,
		retries:   retries,
	}
}

func (s *EmailService) SendEmail(from, subject, body, to string) error {
	var lastErr error

	for attempt := 0; attempt <= s.retries; attempt++ {
		for _, provider := range s.providers {
			id, err := provider.SendEmail(from, subject, body, to)
			if err != nil {
				log.Printf("[%T] Failed: %s\n", provider, err)
				lastErr = err
				continue
			}

			log.Printf("Transmission sent with id [%s] using [%T]\n", id, provider)
			return nil
		}

		if attempt < s.retries {
			log.Printf("All providers failed, retrying in 5 seconds (Attempt %d of %d)\n", attempt+1, s.retries)
			time.Sleep(5 * time.Second)
		}
	}

	log.Printf("All providers failed after %d retries\n", s.retries)
	return lastErr
}
