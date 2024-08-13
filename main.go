package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MartinLupa/go-email-service/config"
	"github.com/MartinLupa/go-email-service/providers"
	"github.com/MartinLupa/go-email-service/service"
	"github.com/joho/godotenv"
)

type EmailPayload struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	To      string `json:"to"`
}

// TODO: add Swagger documentation
// TODO: add testing

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()

	mailgunProvider := providers.NewMailgunProvider(cfg.MailgunDomain, cfg.MailgunAPIKey)
	sparkPostProvider := providers.NewSparkPostProvider(cfg.SparkPostAPIKey)

	providersList := []providers.EmailProvider{mailgunProvider, sparkPostProvider}

	emailService := service.NewEmailService(providersList, 1)

	http.HandleFunc("/send-email", func(w http.ResponseWriter, r *http.Request) {
		var payload EmailPayload

		reqErr := json.NewDecoder(r.Body).Decode(&payload)

		if reqErr != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		err := emailService.SendEmail("sender@example.com", payload.Subject, payload.Body, payload.To)
		if err != nil {
			http.Error(w, "Both providers failed.", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Email sent successfully!"))
	})

	log.Println("Starting server on :" + cfg.PORT)
	log.Fatal(http.ListenAndServe(":"+cfg.PORT, nil))
}
