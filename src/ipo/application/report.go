package application

import "github.com/jorbriib/theIPOGuide/src/ipo/domain"

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

	return s.emailService.Send(report)
}