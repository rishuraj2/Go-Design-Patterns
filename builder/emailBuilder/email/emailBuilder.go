package email

type EmailBuilder struct {
	to          string
	subject     string
	cc          string
	bcc         string
	body        string
	priority    EmailPriority
	attachments string
}

func NewEmailBuilder(to, subject string) *EmailBuilder {
	return &EmailBuilder{
		to: to,
		subject: subject,
		priority: NORMAL,
	}
}

func (eb *EmailBuilder) Cc(cc string) *EmailBuilder {
	eb.cc = cc
	return eb
}

func (eb *EmailBuilder) Bcc(bcc string) *EmailBuilder {
	eb.bcc = bcc
	return eb
}

func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.body = body
	return eb
}

func (eb *EmailBuilder) Priority(priority EmailPriority) *EmailBuilder {
	eb.priority = priority
	return eb
}

func (eb *EmailBuilder) Attachment(attachments string) *EmailBuilder {
	eb.attachments = attachments
	return eb
}

func (eb *EmailBuilder) Build() Email {
	return Email{
		to: eb.to,
		subject: eb.subject,
		cc: eb.cc,
		bcc: eb.bcc,
		body: eb.body,
		priority: eb.priority,
		attachments: eb.attachments,
	}
}
