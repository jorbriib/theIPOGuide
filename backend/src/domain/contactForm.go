package domain

// ContactForm comes from the client
type ContactForm struct {
	name    string
	email string
	message    string
}

// NewReport returns the Report struct
func NewContactForm(name string, email string, message string) ContactForm {
	return ContactForm{name: name, email: email, message: message}
}

// Name returns the name
func (c ContactForm) Name() string {
	return c.name
}

// Email returns the email
func (c ContactForm) Email() string {
	return c.email
}

// Message returns the contact message
func (c ContactForm) Message() string {
	return c.message
}
