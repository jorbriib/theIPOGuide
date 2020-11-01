package infrastructure_test

import (
	"github.com/jorbriib/theIPOGuide/backend/src/infrastructure"
	"github.com/stretchr/testify/assert"
	"net/smtp"
	"testing"
)

func TestNewEmailConfig(t *testing.T) {
	config := infrastructure.NewEmailConfig("", 3000, "", "", "", "")

	assert.IsType(t, infrastructure.EmailConfig{}, config)

}

func TestNewSmtpEmailService(t *testing.T) {
	config := infrastructure.NewEmailConfig("", 3000, "", "", "", "")
	service := infrastructure.NewSmtpEmailService(config, smtp.SendMail)

	assert.IsType(t, infrastructure.SmtpEmailService{}, service)
}

func TestSmtpEmailService_Send(t *testing.T) {
	f, r := mockSend(nil)
	config := infrastructure.NewEmailConfig("http://server.com", 3000, "username", "password", "sender@gmail.com", "default@gmail.com")
	service := infrastructure.NewSmtpEmailService(config, f)

	err := service.Send("email@email.om", "This is the subject", "This is the body")

	assert.Nil(t, err)
	assert.Equal(t, "From: sender@gmail.com\nTo: email@email.om\nSubject: This is the subject\n\nThis is the body", string(r.msg))
	assert.Equal(t, "http://server.com:3000", r.addr)
	assert.Equal(t, "sender@gmail.com", r.from)
	assert.Equal(t, []string{"email@email.om"}, r.to)

	emailAuth := smtp.PlainAuth("", "username", "password", "http://server.com")
	assert.Equal(t, emailAuth, r.auth)
}


func TestSmtpEmailService_Send_WithDefaultSender(t *testing.T) {
	f, r := mockSend(nil)
	config := infrastructure.NewEmailConfig("http://server.com", 3000, "username", "password", "sender@gmail.com", "default@gmail.com")
	service := infrastructure.NewSmtpEmailService(config, f)

	err := service.Send("", "This is the subject", "This is the body")

	assert.Nil(t, err)
	assert.Equal(t, "From: sender@gmail.com\nTo: default@gmail.com\nSubject: This is the subject\n\nThis is the body", string(r.msg))
	assert.Equal(t, "http://server.com:3000", r.addr)
	assert.Equal(t, "sender@gmail.com", r.from)
	assert.Equal(t, []string{"default@gmail.com"}, r.to)

	emailAuth := smtp.PlainAuth("", "username", "password", "http://server.com")
	assert.Equal(t, emailAuth, r.auth)
}

type emailRecorder struct {
	addr string
	auth smtp.Auth
	from string
	to   []string
	msg  []byte
}

func mockSend(errToReturn error) (func(string, smtp.Auth, string, []string, []byte) error, *emailRecorder) {
	r := new(emailRecorder)
	return func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		*r = emailRecorder{addr, a, from, to, msg}
		return errToReturn
	}, r
}
