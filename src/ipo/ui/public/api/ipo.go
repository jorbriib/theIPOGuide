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

type GetIpoJsonResponse struct {
	CompanyName string `json:"companyName"`
	MarketName  string `json:"marketName"`
}

type GetIposJsonResponse []GetIpoJsonResponse

func (c Controller) GetIpos(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := application.NewGetIposQuery()
	response, err := c.service.GetIPOs(query)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	ipos, markets, companies := response.Get()

	jsonResponse := make([]GetIpoJsonResponse, len(ipos))
	for k, ipo := range ipos {
		companyName := ""
		for _, company := range companies{
			if company.Id() == ipo.CompanyId(){
				companyName = company.Name()
				break
			}
		}

		marketName := ""
		for _, market := range markets{
			if market.Id() == ipo.MarketId(){
				marketName = market.Name()
				break
			}
		}
		jsonResponse[k] = GetIpoJsonResponse{CompanyName: companyName, MarketName: marketName}
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(jsonResponse)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
