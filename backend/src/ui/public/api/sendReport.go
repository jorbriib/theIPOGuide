package api

import (
	"encoding/json"
	"github.com/jorbriib/theIPOGuide/backend/src/application"
	"log"
	"net/http"
)

// SendReportController is the struct to send reports
type SendReportController struct {
	service application.SendReportService
}


// NewSendReportController returns a SendReportController struct
func NewSendReportController(service application.SendReportService) SendReportController {
	return SendReportController{service: service}
}

// SendReport sends the report to the email
func (c SendReportController) Run(writer http.ResponseWriter, request *http.Request) {

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
	err := c.service.Run(command)

	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusUnprocessableEntity)
		msg := ErrorMessage{Message: "Forms.ErrorMessages.DefaultErrorMessage"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusAccepted)
}
