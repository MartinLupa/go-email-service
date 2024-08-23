package service

import (
	"time"

	"github.com/MartinLupa/go-email-service/config"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type EmailWorkflowParams struct {
	Config  *config.Config
	From    string
	Subject string
	Body    string
	To      string
}

func SendEmail(ctx workflow.Context, params EmailWorkflowParams) (string, error) {
	var emailID string

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 5,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Second * 5,
			MaximumAttempts:    3,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, SendEmailViaMailgun, params.Config, params.From, params.Subject, params.Body, params.To).Get(ctx, &emailID)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to send email via Mailgun, trying SparkPost", "error", err)
		err = workflow.ExecuteActivity(ctx, SendEmailViaSparkPost, params.Config, params.From, params.Subject, params.Body, params.To).Get(ctx, &emailID)
	}

	return emailID, err
}
