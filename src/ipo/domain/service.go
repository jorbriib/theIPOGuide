package domain

type EmailService interface {
	Send(report Report) error
}
