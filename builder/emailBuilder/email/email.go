package email

import "fmt"

type Email struct {
	to          string
	subject     string
	cc          string
	bcc         string
	body        string
	priority    EmailPriority
	attachments string
}

func (e Email) ToString() {
	fmt.Printf("To: %s\nSubject: %s\nCC: %s\nBCC: %s\nBody: %s\nPriority: %s\nAttachments: %s\n", e.to, e.subject, e.cc, e.bcc, e.body, e.priority.String(), e.attachments)
}
