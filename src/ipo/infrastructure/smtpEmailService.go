package infrastructure

import (
	"fmt"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"net/smtp"
)

// SmtpEmailService is the needed data to send an email
type SmtpEmailService struct {
	emailHost     string
	emailFrom     string
	emailTo       string
	emailPassword string
	emailPort     int
}

// NewSmtpEmailService creates a SmtpEmailService struct
func NewSmtpEmailService(emailHost string, emailFrom string, emailTo string, emailPassword string, emailPort int) SmtpEmailService {
	return SmtpEmailService{emailHost: emailHost, emailFrom: emailFrom, emailTo: emailTo, emailPassword: emailPassword, emailPort: emailPort}
}

// Send sends the report to the email
func (s SmtpEmailService) Send(report domain.Report) error {
	emailAuth := smtp.PlainAuth("", s.emailFrom, s.emailPassword, s.emailHost)
	emailBody := fmt.Sprintf("We have a new report in %s: %s", report.Url(), report.Message())
	subject := "The IPO guide report"

	msg := "From: " + s.emailFrom + "\n" +
		"To: " + s.emailTo + "\n" +
		"Subject: " + subject + "\n\n" +
		emailBody

	addr := fmt.Sprintf("%s:%d", s.emailHost, s.emailPort)

	return smtp.SendMail(addr, emailAuth, s.emailFrom, []string{s.emailTo}, []byte(msg))
}
