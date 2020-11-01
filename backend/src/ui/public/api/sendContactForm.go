package api

import (
	"encoding/json"
	"github.com/jorbriib/theIPOGuide/backend/src/application"
	"log"
	"net/http"
)

// SendContactFormController is the struct to send reports
type SendContactFormController struct {
	service application.SendContactFormService
}

// NewSendContactFormController returns a SendContactFormController struct
func NewSendContactFormController(service application.SendContactFormService) SendContactFormController {
	return SendContactFormController{service: service}
}

// SendContactForm sends the report to the email
func (c SendContactFormController) Run(writer http.ResponseWriter, request *http.Request) {

	name := request.FormValue("name")
	email := request.FormValue("email")
	message := request.FormValue("message")

	if name == "" || email == "" || message == "" {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		msg := ErrorMessage{Message: "Forms.ErrorMessages.SomeInputIsEmpty"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	command := application.SendContactFormCommand{
		Name:    name,
		Email:   email,
		Message: message,
	}
	err := c.service.Run(command)

	if err != nil {
		log.Println(err)

		writer.WriteHeader(http.StatusUnprocessableEntity)
		defaultMessage := "Forms.ErrorMessages.DefaultErrorMessage"

		re, ok := err.(*application.AppError)
		if ok{
			defaultMessage = re.Error()
		}

		msg := ErrorMessage{Message: defaultMessage}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusAccepted)
}
