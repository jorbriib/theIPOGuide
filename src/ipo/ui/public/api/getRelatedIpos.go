package api

import (
	"encoding/json"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"net/http"
)

type GetRelatedIposController struct {
	service application.GetRelatedIposService
}

func NewGetRelatedIposController(service application.GetRelatedIposService) GetRelatedIposController {
	return GetRelatedIposController{service: service}
}

func (c GetRelatedIposController) Run(writer http.ResponseWriter, request *http.Request) {

	query := application.NewGetRelatedIposQuery()
	response, err := c.service.Run(query)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	markets, countries, sectors := response.Get()

	var marketsJsonResponse []*MarketJsonResponse
	for _, market := range markets {
		marketsJsonResponse = append(marketsJsonResponse, &MarketJsonResponse{
			Code:     market.Code(),
			Name:     market.Name(),
			Currency: market.Currency().Name(),
		})
	}

	var countriesJsonResponse []*CountryJsonResponse
	for _, country := range countries {
		countriesJsonResponse = append(countriesJsonResponse, &CountryJsonResponse{
			Code: country.Code(),
			Name: country.Name(),
		})
	}

	var sectorsJsonResponse []*SectorJsonResponse
	for _, sector := range sectors {
		sectorsJsonResponse = append(sectorsJsonResponse, &SectorJsonResponse{
			Alias: sector.Alias(),
			Name:  sector.Name(),
		})
	}

	jsonResponse := &IpoRelations{
		Markets:   marketsJsonResponse,
		Countries: countriesJsonResponse,
		Sectors:   sectorsJsonResponse,
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(jsonResponse)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
