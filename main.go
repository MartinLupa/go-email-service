package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/MartinLupa/go-email-service/config"
	app "github.com/MartinLupa/go-email-service/service"
	"go.temporal.io/sdk/client"
)

type EmailPayload struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	To      string `json:"to"`
}

// TODO: add Swagger documentation
// TODO: add testing
// TODO: add structured error messages
// TODO: add payload validation

func main() {
	// Load environment variables
	cfg := config.LoadConfig()

	// Create the Temporal client object just once per process
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	// Define service endpoint
	http.HandleFunc("/send-email", func(w http.ResponseWriter, r *http.Request) {
		const SendEmailQueue = "SEND_EMAIL_QUEUE"
		var payload EmailPayload
		var result string

		// Parse payload
		reqErr := json.NewDecoder(r.Body).Decode(&payload)

		if reqErr != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Config and start Temporal workflow
		options := client.StartWorkflowOptions{
			ID:        "send-email-workflow",
			TaskQueue: SendEmailQueue,
		}

		input := app.EmailWorkflowParams{
			Config:  cfg,
			From:    "sender@gmail.com",
			Subject: payload.Subject,
			Body:    payload.Body,
			To:      payload.To,
		}

		log.Printf("Starting email sending workflow...\n")

		we, err := c.ExecuteWorkflow(context.Background(), options, app.SendEmail, input)

		if err != nil {
			log.Fatalln("Unable to start the Workflow:", err)
		}

		log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

		err = we.Get(context.Background(), &result)

		if err != nil {
			log.Fatalln("Unable to get Workflow result:", err)
		}

		log.Println(result)

		// Return response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Email sent successfully"))
	})

	// Start server
	log.Println("Starting server on :" + cfg.PORT)
	log.Fatal(http.ListenAndServe(":"+cfg.PORT, nil))
}
