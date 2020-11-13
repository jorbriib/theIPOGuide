package api

import (
	"encoding/json"
	"github.com/jorbriib/theIPOGuide/backend/src/application"
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

	var marketsJsonResponse []*ExtendedMarketJsonResponse
	for _, market := range markets {
		marketsJsonResponse = append(marketsJsonResponse, &ExtendedMarketJsonResponse{
			Code:     market.Code(),
			Name:     market.Name(),
			Currency: market.Currency().Name(),
			Image: market.Image(),
			TotalIpos: market.TotalIpos(),
		})
	}

	var countriesJsonResponse []*ExtendedCountryJsonResponse
	for _, country := range countries {
		countriesJsonResponse = append(countriesJsonResponse, &ExtendedCountryJsonResponse{
			Code: country.Code(),
			Name: country.Name(),
			Image: country.Image(),
			TotalIpos: country.TotalIpos(),
		})
	}

	var sectorsJsonResponse []*ExtendedSectorJsonResponse
	for _, sector := range sectors {
		sectorsJsonResponse = append(sectorsJsonResponse, &ExtendedSectorJsonResponse{
			Alias: sector.Alias(),
			Name:  sector.Name(),
			Image: sector.Image(),
			TotalIpos: sector.TotalIpos(),
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
