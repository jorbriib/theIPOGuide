package domain

type EmailService interface {
	IsValid(email string) bool
	Send(to string, subject string, body string) error
}
