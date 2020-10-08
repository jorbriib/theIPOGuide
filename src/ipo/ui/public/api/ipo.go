package api

import (
	"encoding/json"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"net/http"
)

type Controller struct {
	service application.Service
}

func NewController(service application.Service) Controller {
	return Controller{service: service}
}

type GetIpoResponse struct {
	CompanyName string `json:"companyName"`
	MarketName  string `json:"marketName"`
}

type GetIposResponse []GetIpoResponse

func (c Controller) GetIpos(writer http.ResponseWriter, request *http.Request) {
	query := application.NewGetIposQuery()
	ipos, _ := c.service.GetIPOs(query)

	response := make([]GetIpoResponse, len(ipos))
	for k, ipo := range ipos {
		response[k] = GetIpoResponse{CompanyName: ipo.Company().Name(), MarketName: ipo.Market().Name()}
	}

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
