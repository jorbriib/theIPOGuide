package api

import (
	"encoding/json"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"net/http"
	"strconv"
	"strings"
)

type GetIposController struct {
	service application.GetIposService
}

func NewGetIposController(service application.GetIposService) GetIposController {
	return GetIposController{service: service}
}

func (c GetIposController) Run(writer http.ResponseWriter, request *http.Request) {

	marketStringCodes := request.URL.Query().Get("markets")
	countryStringCodes := request.URL.Query().Get("countries")
	sectorStringAliases := request.URL.Query().Get("sectors")
	page := request.URL.Query().Get("page")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	query := application.NewGetIposQuery(
		strings.Split(marketStringCodes, ","),
		strings.Split(countryStringCodes, ","),
		strings.Split(sectorStringAliases, ","),
		uint(pageInt),
	)
	response, err := c.service.Run(query)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	total, ipos, markets, companies := response.Get()

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

	output := GetIposJsonResponse{
		Total: total,
		List:  list,
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(output)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
