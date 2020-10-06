package api

import (
	"fmt"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"net/http"
)

type Controller struct {
	service application.Service
}

func NewController(service application.Service) Controller {
	return Controller{service: service} 
}

func (c Controller) GetIpos(writer http.ResponseWriter, request *http.Request) {
	query := application.NewGetIposQuery()
	response, _ := c.service.GetIPOs(query)

	fmt.Fprint(writer, response)
}

