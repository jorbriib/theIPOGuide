package application_test

import (
	"fmt"
	"github.com/jorbriib/theIPOGuide/backend/src/application"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewReportService(t *testing.T) {
	assertion := assert.New(t)
	es := EmailServiceMock{}
	service := application.NewSendReportService(es)
	assertion.NotNil(service)
	assertion.IsType(application.SendReportService{}, service)
}

func TestReportService_SendReport(t *testing.T) {
	assertion := assert.New(t)


	es := EmailServiceMock{}

	expectedReport := domain.NewReport("http://www.url.com", "This is a message")

	emailBody := fmt.Sprintf("We have a new report in %s: %s", expectedReport.Url(), expectedReport.Message())
	subject := "The IPO guide report"

	es.On("Send", "", subject, emailBody).Return(nil)

	service := application.NewSendReportService(es)

	command := application.SendReportCommand{
		Url:     "http://www.url.com",
		Message: "This is a message",
	}
	err := service.SendReport(command)

	assertion.Nil(err)
}

type EmailServiceMock struct {
	mock.Mock
}

func (s EmailServiceMock) Send(to string, subject string, body string) error {
	args := s.Called(to, subject, body)
	return args.Error(0)
}