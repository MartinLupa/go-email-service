package providers

type EmailProvider interface {
	SendEmail(from, subject, body, to string) (string, error)
}
