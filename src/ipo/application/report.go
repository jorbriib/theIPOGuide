package application

import (
	"fmt"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
)

// ReportService is the service to send reports
type ReportService struct {
	emailService     domain.EmailService
}

func NewReportService(service domain.EmailService) ReportService{
	return ReportService{service}
}

// SendReportCommand is the command to send a report
type SendReportCommand struct {
	Url string
	Message string
}

// SendReport sends the report to the email
func (s *ReportService) SendReport(command SendReportCommand) error {
	report := domain.NewReport(command.Url, command.Message)

	emailBody := fmt.Sprintf("We have a new report in %s: %s", report.Url(), report.Message())
	subject := "The IPO guide report"

	err := s.emailService.Send("", subject, emailBody)
	return err
}