package application_test

import (
	"errors"
	. "github.com/jorbriib/theIPOGuide/backend/src/application"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSearchByTextService(t *testing.T) {
	ir := IpoRepositoryMock{}
	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}
	ctr := CountryRepositoryMock{}
	sr := SectorRepositoryMock{}
	service := NewSearchByTextService(ir, mr, cr, ctr, sr)
	assert.NotNil(t, service)
}

func TestNewSearchByTextQuery(t *testing.T) {
	query := NewSearchByTextQuery("text")
	assert.NotNil(t, query)
}

func TestSearchByTextService_Run_FailsWhenIpoRepositorySearchByTextReturnsError(t *testing.T) {
	assertion := assert.New(t)

	ir := IpoRepositoryMock{}
	ir.On("SearchByText", "text").Return([]domain.Ipo{}, errors.New("repository error"))

	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}
	ctr := CountryRepositoryMock{}
	sr := SectorRepositoryMock{}
	service := NewSearchByTextService(ir, mr, cr, ctr, sr)

	query := NewSearchByTextQuery("text")

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestSearchByTextService_Run_FailsWhenMarketRepositoryFindByIdsReturnsError(t *testing.T) {
	assertion := assert.New(t)

	ir := IpoRepositoryMock{}
	ir.On("SearchByText", "text").Return([]domain.Ipo{}, nil)

	mr := MarketRepositoryMock{}
	mr.On("FindByIds", []domain.MarketId{}).Return([]domain.Market{}, errors.New("repository error"))

	cr := CompanyRepositoryMock{}
	ctr := CountryRepositoryMock{}
	sr := SectorRepositoryMock{}
	service := NewSearchByTextService(ir, mr, cr, ctr, sr)

	query := NewSearchByTextQuery("text")

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestSearchByTextService_Run_FailsWhenCompanyRepositoryFindByIdsReturnsError(t *testing.T) {
	assertion := assert.New(t)

	ir := IpoRepositoryMock{}
	ir.On("SearchByText", "text").Return([]domain.Ipo{}, nil)

	mr := MarketRepositoryMock{}
	mr.On("FindByIds", []domain.MarketId{}).Return([]domain.Market{}, nil)

	cr := CompanyRepositoryMock{}
	cr.On("FindByIds", []domain.CompanyId{}).Return([]domain.Company{}, errors.New("repository error"))
	ctr := CountryRepositoryMock{}
	sr := SectorRepositoryMock{}
	service := NewSearchByTextService(ir, mr, cr, ctr, sr)

	query := NewSearchByTextQuery("text")

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}


func TestSearchByTextService_Run(t *testing.T) {
	assertion := assert.New(t)

	expectedIpoReturn := make([]domain.Ipo, 3)
	expectedIpoReturn[0] = domain.HydrateIpo("1-ipo-id", "1-alias", "intro", "1-market-id", "1-company-id", 0, 0, 0, nil)
	expectedIpoReturn[1] = domain.HydrateIpo("2-ipo-id", "2-alias", "intro", "2-market-id", "2-company-id", 0, 0, 0, nil)
	expectedIpoReturn[2] = domain.HydrateIpo("3-ipo-id", "3-alias", "intro","2-market-id", "3-company-id", 0, 0, 0, nil)

	expectedMarketReturn := make([]domain.Market, 2)
	expectedMarketReturn[0] = domain.HydrateMarket("1-market-id", "market-code-1", "Market 1", domain.HydrateCurrency("", "", ""))
	expectedMarketReturn[1] = domain.HydrateMarket("2-market-id", "market-code-2", "Market 2", domain.HydrateCurrency("", "", ""))

	expectedCompanyReturn := make([]domain.Company, 3)
	expectedCompanyReturn[0] = domain.HydrateCompany(
		"1-company-id", "", "",
		domain.HydrateSector("1-sector-id", "sector", "Sector"),
		domain.HydrateIndustry("1-industry-id", "industry", "Industry"), "",
		domain.HydrateCountry("1-country-id", "es", "Spain"), "",
		"", "", 0, "", "",
		"", "", "", "", 2000, "", "", "", "", "")
	expectedCompanyReturn[1] = domain.HydrateCompany(
		"2-company-id", "", "",
		domain.HydrateSector("1-sector-id", "sector", "Sector"),
		domain.HydrateIndustry("1-industry-id", "industry", "Industry"), "",
		domain.HydrateCountry("1-country-id", "es", "Spain"), "",
		"", "", 0, "", "",
		"", "", "", "", 2000, "", "", "", "", "")
	expectedCompanyReturn[2] = domain.HydrateCompany(
		"3-company-id", "", "",
		domain.HydrateSector("2-sector-id", "2-sector", "Sector 2" ),
		domain.HydrateIndustry("2-industry-id", "2-industry", "Industry 2"), "",
		domain.HydrateCountry("2-country-id", "us", "USA"), "",
		"", "", 0, "", "",
		"", "", "", "", 2000, "", "", "", "", "")

	expectedMarketIdInput := make([]domain.MarketId, 2)
	expectedMarketIdInput[0] = "1-market-id"
	expectedMarketIdInput[1] = "2-market-id"

	expectedCompanyIdInput := make([]domain.CompanyId, 3)
	expectedCompanyIdInput[0] = "1-company-id"
	expectedCompanyIdInput[1] = "2-company-id"
	expectedCompanyIdInput[2] = "3-company-id"

	ir := IpoRepositoryMock{}
	ir.On("SearchByText", "text").Return(expectedIpoReturn, nil)

	mr := MarketRepositoryMock{}
	mr.On("FindByIds", expectedMarketIdInput).Return(expectedMarketReturn, nil)

	cr := CompanyRepositoryMock{}
	cr.On("FindByIds", expectedCompanyIdInput).Return(expectedCompanyReturn, nil)
	ctr := CountryRepositoryMock{}
	sr := SectorRepositoryMock{}
	service := NewSearchByTextService(ir, mr, cr, ctr, sr)

	query := NewSearchByTextQuery("text")

	response, err := service.Run(query)

	assertion.NotNil(response)

	ipos, markets, companies := response.Get()
	assertion.Equal(expectedIpoReturn, ipos)
	assertion.Equal(expectedMarketReturn, markets)
	assertion.Equal(expectedCompanyReturn, companies)
	assertion.Nil(err)
}