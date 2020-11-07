package application_test

import (
	"errors"
	. "github.com/jorbriib/theIPOGuide/backend/src/application"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewSendContactFormService(t *testing.T) {
	es := EmailServiceMock{}
	service := NewSendContactFormService(es)
	assert.NotNil(t, service)
}

func TestSendContactFormService_Run_FailsWhenEmailIsInvalid(t *testing.T) {
	es := EmailServiceMock{}
	es.On("IsValid", "invalid-email").Return(false)

	service := NewSendContactFormService(es)
	command := SendContactFormCommand{
		Name:    "John",
		Email:   "invalid-email",
		Message: "message",
	}
	err := service.Run(command)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid email", err.Error())
}

func TestSendContactFormService_Run_FailsWhenNameTooShort(t *testing.T) {
	es := EmailServiceMock{}
	es.On("IsValid", "john@email.com").Return(true)

	service := NewSendContactFormService(es)
	command := SendContactFormCommand{
		Name:    "",
		Email:   "john@email.com",
		Message: "message",
	}
	err := service.Run(command)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid name, minimum length is 3 characters", err.Error())
}

func TestSendContactFormService_Run_FailsWhenMessageTooShort(t *testing.T) {
	es := EmailServiceMock{}
	es.On("IsValid", "john@email.com").Return(true)

	service := NewSendContactFormService(es)
	command := SendContactFormCommand{
		Name:    "John",
		Email:   "john@email.com",
		Message: "",
	}
	err := service.Run(command)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid message, minimum length is 10 characters", err.Error())
}

func TestSendContactFormService_Run_FailsWhenEmailServiceSendReturnsError(t *testing.T) {
	es := EmailServiceMock{}
	es.On("IsValid", "john@email.com").Return(true)
	es.On("Send", "", "The IPO guide contact form", "We have a new message of John (john@email.com): message").Return(errors.New("service error"))

	service := NewSendContactFormService(es)
	command := SendContactFormCommand{
		Name:    "John",
		Email:   "john@email.com",
		Message: "message",
	}
	err := service.Run(command)

	assert.NotNil(t, err)
}

func TestSendContactFormService_Run(t *testing.T) {
	es := EmailServiceMock{}
	es.On("IsValid", "john@email.com").Return(true)
	es.On("Send", "", "The IPO guide contact form", "We have a new message of John (john@email.com): message").Return(nil)

	service := NewSendContactFormService(es)
	command := SendContactFormCommand{
		Name:    "John",
		Email:   "john@email.com",
		Message: "message",
	}
	err := service.Run(command)

	assert.Nil(t, err)
}

type EmailServiceMock struct {
	mock.Mock
}

func (s EmailServiceMock) Send(to string, subject string, body string) error {
	args := s.Called(to, subject, body)
	return args.Error(0)
}

func (s EmailServiceMock) IsValid(email string) bool {
	args := s.Called(email)
	return args.Get(0).(bool)
}

