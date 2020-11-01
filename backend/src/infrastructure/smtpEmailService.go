package infrastructure

import (
	"fmt"
	"net"
	"net/smtp"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type EmailConfig struct {
	serverHost      string
	serverPort      int
	username        string
	password        string
	sender          string
	defaultReceiver string
}

// smtpEmailService is the needed data to send an email
type SmtpEmailService struct {
	conf EmailConfig
	send func(string, smtp.Auth, string, []string, []byte) error
}

// NewSmtpEmailService creates a smtpEmailService struct
func NewEmailConfig(serverHost string, serverPort int, username string, password string, sender string, defaultReceiver string) EmailConfig {
	return EmailConfig{serverHost: serverHost, serverPort: serverPort, username: username, password: password, sender: sender, defaultReceiver: defaultReceiver}
}

// NewSmtpEmailService creates a smtpEmailService struct
func NewSmtpEmailService(conf EmailConfig, send func(string, smtp.Auth, string, []string, []byte) error) SmtpEmailService {
	return SmtpEmailService{conf, send}
}

// Send sends the report to the email
func (s SmtpEmailService) Send(to string, subject string, body string) error {
	if to == "" {
		to = s.conf.defaultReceiver
	}
	emailAuth := smtp.PlainAuth("", s.conf.username, s.conf.password, s.conf.serverHost)

	msg := "From: " + s.conf.sender + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	addr := fmt.Sprintf("%s:%d", s.conf.serverHost, s.conf.serverPort)

	return s.send(addr, emailAuth, s.conf.sender, []string{to}, []byte(msg))
}

// IsValid checks if email is valid
func (s SmtpEmailService) IsValid(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	if !emailRegex.MatchString(email) {
		return false
	}
	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}
