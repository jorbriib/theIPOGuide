package api

import (
	"encoding/json"
	"github.com/golossus/routing"
	"github.com/jorbriib/theIPOGuide/backend/src/application"
	"net/http"
)

type GetIpoController struct {
	service application.GetIpoService
}

func NewGetIpoController(service application.GetIpoService) GetIpoController {
	return GetIpoController{service: service}
}

func (c GetIpoController) Run(writer http.ResponseWriter, request *http.Request) {

	bag := routing.GetURLParameters(request)
	alias, err := bag.GetByName("alias")
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
	}

	query := application.NewGetIpoQuery(alias)
	response, err := c.service.Run(query)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	ipo, market, company := response.Get()

	companyJsonResponse := &CompanyViewJsonResponse{}
	if company != nil {
		companyJsonResponse.Symbol = company.Symbol()
		companyJsonResponse.Name = company.Name()
		companyJsonResponse.Description = company.Description()
		companyJsonResponse.Employees = company.Employees()

		companyJsonResponse.Sector = company.Sector().Name()
		companyJsonResponse.Industry = company.Industry().Name()

		companyJsonResponse.Address = company.Address()
		companyJsonResponse.Country = company.Country().Name()

		companyJsonResponse.Phone = company.Phone()
		companyJsonResponse.Email = company.Email()
		companyJsonResponse.Website = company.Website()

		companyJsonResponse.Facebook = company.Facebook()
		companyJsonResponse.Twitter = company.Twitter()
		companyJsonResponse.Linkedin = company.Linkedin()
		companyJsonResponse.Pinterest = company.Pinterest()
		companyJsonResponse.Instagram = company.Instagram()

		companyJsonResponse.Ceo = company.Ceo()
		companyJsonResponse.Founded = company.Founded()
		companyJsonResponse.FiscalYearEnd = company.FiscalYearEnd()
		companyJsonResponse.IpoUrl = company.IpoUrl()
		companyJsonResponse.ExchangeCommissionUrl = company.ExchangeCommissionUrl()

		companyJsonResponse.Logo = company.LogoUrl()
	}

	marketJsonResponse := &MarketJsonResponse{}
	var priceFrom string
	var priceTo string
	if market != nil {
		marketJsonResponse.Code = market.Code()
		marketJsonResponse.Name = market.Name()
		marketJsonResponse.Currency = market.Currency().Code()
		priceFrom = market.Currency().DisplayFromCents(ipo.PriceCentsFrom())
		if ipo.PriceCentsTo() != 0 {
			priceTo = market.Currency().DisplayFromCents(ipo.PriceCentsTo())
		}
	}

	jsonResponse := IpoViewJsonResponse{
		ipo.Alias(),
		ipo.Intro(),
		companyJsonResponse,
		marketJsonResponse,
		priceFrom,
		priceTo,
		ipo.Shares(),
		ipo.ExpectedDate().String(),
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(jsonResponse)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
