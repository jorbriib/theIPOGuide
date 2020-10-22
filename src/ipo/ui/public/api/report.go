package api

import (
	"encoding/json"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"log"
	"net/http"
)

// ReportController is the struct to send reports
type ReportController struct {
	service application.ReportService
}

// ErrorMessage is the error message to send to the clients
type ErrorMessage struct {
	Message string `json:"message"`
}

// NewReportController returns a ReportController struct
func NewReportController(service application.ReportService) ReportController {
	return ReportController{service: service}
}

// SendFeedback sends the report to the email
func (c ReportController) SendReport(writer http.ResponseWriter, request *http.Request) {

	url := request.FormValue("url")
	message := request.FormValue("message")


	if url == "" || message == "" {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		msg := ErrorMessage{Message: "Forms.ErrorMessages.MessageOrUrlAreEmpty"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	command := application.SendReportCommand{
		Url:     url,
		Message: message,
	}
	err := c.service.SendReport(command)

	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusUnprocessableEntity)
		msg := ErrorMessage{Message: "Forms.ErrorMessages.DefaultErrorMessage"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusAccepted)
}
