package api

import (
	"encoding/json"
	"github.com/golossus/routing"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"net/http"
)

type Controller struct {
	service application.Service
}

func NewController(service application.Service) Controller {
	return Controller{service: service}
}

type IpoJsonResponse struct {
	Alias        string               `json:"alias"`
	Company      *CompanyJsonResponse `json:"company"`
	Market       *MarketJsonResponse  `json:"market"`
	PriceFrom    string               `json:"priceFrom"`
	PriceTo      string               `json:"priceTo"`
	ExpectedDate string               `json:"expectedDate"`
}

type CompanyJsonResponse struct {
	Symbol  string `json:"symbol"`
	Name    string `json:"name"`
	Sector  string `json:"sector"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
}

type MarketJsonResponse struct {
	Name string `json:"name"`
}

type GetIposJsonResponse []IpoJsonResponse

func (c Controller) GetIpos(writer http.ResponseWriter, request *http.Request) {

	query := application.NewGetIposQuery()
	response, err := c.service.GetIPOs(query)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ipos, markets, companies := response.Get()

	jsonResponse := make([]IpoJsonResponse, len(ipos))
	for k, ipo := range ipos {
		companyJsonResponse := &CompanyJsonResponse{}
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
				marketJsonResponse.Name = market.Name()
				priceFrom = market.Currency().DisplayFromCents(ipo.PriceCentsFrom())
				priceTo = market.Currency().DisplayFromCents(ipo.PriceCentsTo())
				break
			}
		}

		jsonResponse[k] = IpoJsonResponse{
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

func (c Controller) GetIpo(writer http.ResponseWriter, request *http.Request) {

	bag := routing.GetURLParameters(request)
	alias, err := bag.GetByName("alias")
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
	}

	query := application.NewGetIpoQuery(alias)
	response, err := c.service.GetIPO(query)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	ipo, market, company := response.Get()

	companyJsonResponse := &CompanyJsonResponse{}
	if company != nil {
		companyJsonResponse.Symbol = company.Symbol()
		companyJsonResponse.Name = company.Name()
		companyJsonResponse.Sector = company.Sector().Name()
		companyJsonResponse.Country = company.Country().Name()
		companyJsonResponse.Logo = company.LogoUrl()
	}

	marketJsonResponse := &MarketJsonResponse{}
	var priceFrom string
	var priceTo string
	if market != nil {
		marketJsonResponse.Name = market.Name()
		priceFrom = market.Currency().DisplayFromCents(ipo.PriceCentsFrom())
		priceTo = market.Currency().DisplayFromCents(ipo.PriceCentsTo())
	}

	jsonResponse := IpoJsonResponse{
		ipo.Alias(),
		companyJsonResponse,
		marketJsonResponse,
		priceFrom,
		priceTo,
		ipo.ExpectedDate().String(),
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(jsonResponse)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}