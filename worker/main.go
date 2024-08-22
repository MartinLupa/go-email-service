package main

import (
	"log"

	"github.com/MartinLupa/go-email-service/app"
	"github.com/MartinLupa/go-email-service/config"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	cfg := config.LoadConfig()

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, cfg.QueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(app.SendEmail)
	w.RegisterActivity(app.SendEmailViaMailgun)
	w.RegisterActivity(app.SendEmailViaSparkPost)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
