package api

import (
	"fmt"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"net/http"
)

type Controller struct {
	handler application.Handler
}

func NewController(handler application.Handler) Controller {
	return Controller{handler: handler}
}

func (c Controller) GetIpos(writer http.ResponseWriter, request *http.Request) {
	query := application.NewGetIposQuery()
	response := c.handler.GetIPOs(query)

	fmt.Fprint(writer, response)
}

