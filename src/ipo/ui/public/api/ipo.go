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

type IpoListJsonResponse struct {
	Alias        string                   `json:"alias"`
	Company      *CompanyListJsonResponse `json:"company"`
	Market       *MarketJsonResponse      `json:"market"`
	PriceFrom    string                   `json:"priceFrom"`
	PriceTo      string                   `json:"priceTo"`
	ExpectedDate string                   `json:"expectedDate"`
}

type CompanyListJsonResponse struct {
	Symbol  string `json:"symbol"`
	Name    string `json:"name"`
	Sector  string `json:"sector"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
}

type MarketJsonResponse struct {
	Name string `json:"name"`
}

type GetIposJsonResponse []IpoListJsonResponse

func (c Controller) GetIpos(writer http.ResponseWriter, request *http.Request) {

	query := application.NewGetIposQuery()
	response, err := c.service.GetIPOs(query)
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
				marketJsonResponse.Name = market.Name()
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

type IpoViewJsonResponse struct {
	Alias        string                   `json:"alias"`
	Company      *CompanyViewJsonResponse `json:"company"`
	Market       *MarketJsonResponse      `json:"market"`
	PriceFrom    string                   `json:"priceFrom"`
	PriceTo      string                   `json:"priceTo"`
	Shares       uint32                   `json:"shares"`
	ExpectedDate string                   `json:"expectedDate"`
}

type CompanyViewJsonResponse struct {
	Symbol                string `json:"symbol"`
	Name                  string `json:"name"`
	Sector                string `json:"sector"`
	Industry              string `json:"industry"`
	Address               string `json:"address"`
	Country               string `json:"country"`
	Phone                 string `json:"phone"`
	Email                 string `json:"email"`
	Website               string `json:"website"`
	Description           string `json:"description"`
	Facebook              string `json:"facebook"`
	Twitter               string `json:"twitter"`
	Linkedin              string `json:"linkedin"`
	Pinterest             string `json:"pinterest"`
	Instagram             string `json:"instagram"`
	Employees             uint32 `json:"employees"`
	Founded               uint16 `json:"founded"`
	Ceo                   string `json:"ceo"`
	FiscalYearEnd         string `json:"fiscalYearEnd"`
	IpoUrl                string `json:"ipoUrl"`
	ExchangeCommissionUrl string `json:"exchangeCommissionUrl"`
	Logo                  string `json:"logo"`
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
		marketJsonResponse.Name = market.Name()
		priceFrom = market.Currency().DisplayFromCents(ipo.PriceCentsFrom())
		if ipo.PriceCentsTo() != 0 {
			priceTo = market.Currency().DisplayFromCents(ipo.PriceCentsTo())
		}
	}

	jsonResponse := IpoViewJsonResponse{
		ipo.Alias(),
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
