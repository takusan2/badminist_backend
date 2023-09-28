package utils

import (
	"fmt"

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

	sent, err := e.client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(sent.Id)
}
