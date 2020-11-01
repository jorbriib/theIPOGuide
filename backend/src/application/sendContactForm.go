package application

import (
	"fmt"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
)

// SendContactFormService is the service to send reports
type SendContactFormService struct {
	emailService domain.EmailService
}

func NewSendContactFormService(service domain.EmailService) SendContactFormService {
	return SendContactFormService{service}
}

// SendContactFormCommand is the command to send a report
type SendContactFormCommand struct {
	Name    string
	Email   string
	Message string
}

// Run sends the report to the email
func (s *SendContactFormService) Run(command SendContactFormCommand) error {

	isValid := s.emailService.IsValid(command.Email)
	if !isValid {
		return NewAppError("Invalid email")
	}

	if len(command.Name) == 0 {
		return NewAppError("Invalid name, minimum length is 3 characters")
	}

	if len(command.Message) == 0 {
		return NewAppError("Invalid message, minimum length is 10 characters")
	}

	report := domain.NewContactForm(command.Name, command.Email, command.Message)

	emailBody := fmt.Sprintf("We have a new message of %s (%s): %s", report.Name(), report.Email(), report.Message())
	subject := "The IPO guide contact form"

	err := s.emailService.Send("", subject, emailBody)
	return err
}
