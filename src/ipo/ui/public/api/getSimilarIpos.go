package api

import (
	"encoding/json"
	"github.com/golossus/routing"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"net/http"
)

type GetSimilarIposController struct {
	service application.GetSimilarIposService
}

func NewGetSimilarIposController(service application.GetSimilarIposService) GetSimilarIposController {
	return GetSimilarIposController{service: service}
}

func (c GetSimilarIposController) Run(writer http.ResponseWriter, request *http.Request) {

	bag := routing.GetURLParameters(request)
	alias, err := bag.GetByName("alias")
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
	}

	query := application.NewGetSimilarIposQuery(alias)
	response, err := c.service.Run(query)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ipos, markets, companies := response.Get()

	jsonResponse := make([]IpoListJsonResponse, len(ipos))
	for k, ipo := range ipos {
		companyJsonResponse := &CompanyListJsonResponse{}
		for _, company := range companies {
			if company.Id() == ipo.CompanyId() {
				companyJsonResponse.Symbol = company.Symbol()
				companyJsonResponse.Name = company.Name()
				companyJsonResponse.Sector = company.Sector().Name()
				companyJsonResponse.Country = company.Country().Name()
				companyJsonResponse.Logo = company.LogoUrl()
				break
			}
		}

		var priceFrom string
		var priceTo string
		marketJsonResponse := &MarketJsonResponse{}
		for _, market := range markets {
			if market.Id() == ipo.MarketId() {
				marketJsonResponse.Code = market.Code()
				marketJsonResponse.Name = market.Name()
				marketJsonResponse.Currency = market.Currency().Code()
				priceFrom = market.Currency().DisplayFromCents(ipo.PriceCentsFrom())
				if ipo.PriceCentsTo() != 0 {
					priceTo = market.Currency().DisplayFromCents(ipo.PriceCentsTo())
				}
				break
			}
		}

		jsonResponse[k] = IpoListJsonResponse{
			ipo.Alias(),
			companyJsonResponse,
			marketJsonResponse,
			priceFrom,
			priceTo,
			ipo.ExpectedDate().String(),
		}
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(jsonResponse)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
