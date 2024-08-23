package main

import (
	"log"

	"github.com/MartinLupa/go-email-service/config"
	"github.com/MartinLupa/go-email-service/service"
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
	w.RegisterWorkflow(service.SendEmail)
	w.RegisterActivity(service.SendEmailViaMailgun)
	w.RegisterActivity(service.SendEmailViaSparkPost)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
