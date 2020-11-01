package application

import (
	"fmt"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
)

// ReportService is the service to send reports
type SendReportService struct {
	emailService     domain.EmailService
}

func NewSendReportService(service domain.EmailService) SendReportService{
	return SendReportService{service}
}

// SendReportCommand is the command to send a report
type SendReportCommand struct {
	Url string
	Message string
}

// Run sends the report to the email
func (s *SendReportService) Run(command SendReportCommand) error {
	report := domain.NewReport(command.Url, command.Message)

	emailBody := fmt.Sprintf("We have a new report in %s: %s", report.Url(), report.Message())
	subject := "The IPO guide report"

	err := s.emailService.Send("", subject, emailBody)
	return err
}