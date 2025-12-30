package email

import (
	"context"
	"fmt"

	"github.com/taber-miyauchi/notifier-core"
)

// EmailNotifier sends notifications via SMTP.
type EmailNotifier struct {
	SMTPHost string
	SMTPPort int
	From     string
}

// NewEmailNotifier creates an EmailNotifier with default SMTP port.
func NewEmailNotifier(host, from string) *EmailNotifier {
	return &EmailNotifier{
		SMTPHost: host,
		SMTPPort: 587,
		From:     from,
	}
}

// Send delivers a message via email.
// Implements the core.Notifier interface.
func (e *EmailNotifier) Send(ctx context.Context, msg core.Message) error {
	if msg.Recipient == "" {
		return fmt.Errorf("recipient is required")
	}

	fmt.Printf("[EMAIL] From: %s, To: %s, Subject: %s, Priority: %d\n",
		e.From, msg.Recipient, msg.Subject, msg.Priority)

	return nil
}
