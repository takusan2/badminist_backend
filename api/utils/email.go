package utils

import (
	"github.com/resendlabs/resend-go"
)

type EmailServer interface {
	SendEmail(
		to []string,
		subject string,
		html string,
	)
}

type emailServer struct {
	client *resend.Client
}

func NewEmailServer(client *resend.Client) EmailServer {
	return &emailServer{
		client: client,
	}
}

func (e *emailServer) SendEmail(
	to []string,
	subject string,
	html string,
) {
	params := &resend.SendEmailRequest{
		From:    "noreply@badminist.com",
		To:      to,
		Subject: subject,
		Html:    html,
	}

	if _, err := e.client.Emails.Send(params); err != nil {
		return
	}
}
