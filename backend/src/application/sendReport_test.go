package application_test

import (
	"errors"
	"fmt"
	"github.com/jorbriib/theIPOGuide/backend/src/application"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReportService(t *testing.T) {
	assertion := assert.New(t)
	es := EmailServiceMock{}
	service := application.NewSendReportService(es)
	assertion.NotNil(service)
	assertion.IsType(application.SendReportService{}, service)
}


func TestSendReportService_Run_FailsWhenEmailServiceSendReturnsError(t *testing.T) {
	assertion := assert.New(t)

	es := EmailServiceMock{}

	expectedReport := domain.NewReport("http://www.url.com", "This is a message")

	emailBody := fmt.Sprintf("We have a new report in %s: %s", expectedReport.Url(), expectedReport.Message())
	subject := "The IPO guide report"

	es.On("Send", "", subject, emailBody).Return(errors.New("Service error"))

	service := application.NewSendReportService(es)

	command := application.SendReportCommand{
		Url:     "http://www.url.com",
		Message: "This is a message",
	}
	err := service.Run(command)

	assertion.NotNil(err)
}


func TestSendReportService_Run(t *testing.T) {
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
	err := service.Run(command)

	assertion.Nil(err)
}
