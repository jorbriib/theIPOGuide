package api

import (
	"encoding/json"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"net/http"
)

type SearchByTextController struct {
	service application.SearchByTextService
}

func NewSearchByTextController(service application.SearchByTextService) SearchByTextController {
	return SearchByTextController{service: service}
}

func (c SearchByTextController) Run(writer http.ResponseWriter, request *http.Request) {

	text := request.URL.Query().Get("text")
	if text == "" {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if len(text) < 3{
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	query := application.NewSearchByTextQuery(text)
	response, err := c.service.Run(query)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ipos, markets, companies := response.Get()

	list := make([]IpoListJsonResponse, len(ipos))
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

		list[k] = IpoListJsonResponse{
			ipo.Alias(),
			companyJsonResponse,
			marketJsonResponse,
			priceFrom,
			priceTo,
			ipo.ExpectedDate().String(),
		}
	}

	output := SearchByTextJsonResponse{
		List: list,
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(output)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
