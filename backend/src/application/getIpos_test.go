package application_test

import (
	"errors"
	. "github.com/jorbriib/theIPOGuide/backend/src/application"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGetIposService(t *testing.T) {
	r := IpoRepositoryMock{}
	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}
	ctr := CountryRepositoryMock{}
	sr := SectorRepositoryMock{}
	service := NewGetIposService(r, mr, cr, ctr, sr)
	assert.NotNil(t, service)
}

func TestNewGetIposQuery(t *testing.T) {

	query := NewGetIposQuery([]string{"market-code-1"}, []string{"country-code-1"}, []string{"sector-alias-1"}, 0)
	assert.NotNil(t, query)
}

func TestService_GetIPOs_FailsWhenMarketRepositoryFindByCodesReturnsError(t *testing.T) {
	assertion := assert.New(t)

	r := IpoRepositoryMock{}
	mr := MarketRepositoryMock{}
	mr.On("FindByCodes", []string{"market-code-1"}).Return([]domain.Market{}, errors.New("repository error"))

	cr := CompanyRepositoryMock{}
	ctr := CountryRepositoryMock{}
	sr := SectorRepositoryMock{}
	service := NewGetIposService(r, mr, cr, ctr, sr)
	query := NewGetIposQuery([]string{"market-code-1"}, []string{"country-code-1"}, []string{"sector-alias-1"}, 0)

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs_FailsWhenCountryRepositoryFindByCodesReturnsError(t *testing.T) {
	assertion := assert.New(t)

	r := IpoRepositoryMock{}
	mr := MarketRepositoryMock{}
	mr.On("FindByCodes", []string{"market-code-1"}).Return([]domain.Market{}, nil)

	cr := CompanyRepositoryMock{}

	ctr := CountryRepositoryMock{}
	ctr.On("FindByCodes", []string{"country-code-1"}).Return([]domain.Country{}, errors.New("repository error"))
	sr := SectorRepositoryMock{}
	service := NewGetIposService(r, mr, cr, ctr, sr)
	query := NewGetIposQuery([]string{"market-code-1"}, []string{"country-code-1"}, []string{"sector-alias-1"}, 0)

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs_FailsWhenSectorRepositoryFindByAliasesReturnsError(t *testing.T) {
	assertion := assert.New(t)

	r := IpoRepositoryMock{}
	mr := MarketRepositoryMock{}
	mr.On("FindByCodes", []string{"market-code-1"}).Return([]domain.Market{}, nil)

	cr := CompanyRepositoryMock{}

	ctr := CountryRepositoryMock{}
	ctr.On("FindByCodes", []string{"country-code-1"}).Return([]domain.Country{}, nil)

	sr := SectorRepositoryMock{}
	sr.On("FindByAliases", []string{"sector-alias-1"}).Return([]domain.Sector{}, errors.New("repository error"))
	service := NewGetIposService(r, mr, cr, ctr, sr)
	query := NewGetIposQuery([]string{"market-code-1"}, []string{"country-code-1"}, []string{"sector-alias-1"}, 0)

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs_FailsWhenIpoRepositoryFindReturnsError(t *testing.T) {
	assertion := assert.New(t)

	r := IpoRepositoryMock{}
	r.On("Find", []domain.MarketId{}, []domain.CountryId{}, []domain.SectorId{}, []domain.IndustryId{}, []domain.IpoId{}, "", uint(0), uint(20)).Return([]domain.Ipo{}, errors.New("repository error"))
	mr := MarketRepositoryMock{}
	mr.On("FindByCodes", []string{"market-code-1"}).Return([]domain.Market{}, nil)

	cr := CompanyRepositoryMock{}

	ctr := CountryRepositoryMock{}
	ctr.On("FindByCodes", []string{"country-code-1"}).Return([]domain.Country{}, nil)

	sr := SectorRepositoryMock{}
	sr.On("FindByAliases", []string{"sector-alias-1"}).Return([]domain.Sector{}, nil)
	service := NewGetIposService(r, mr, cr, ctr, sr)
	query := NewGetIposQuery([]string{"market-code-1"}, []string{"country-code-1"}, []string{"sector-alias-1"}, 0)

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs_FailsWhenMarketRepositoryFindByIdsReturnsError(t *testing.T) {
	assertion := assert.New(t)

	r := IpoRepositoryMock{}
	r.On("Find", []domain.MarketId{}, []domain.CountryId{}, []domain.SectorId{}, []domain.IndustryId{}, []domain.IpoId{}, "", uint(0), uint(20)).Return([]domain.Ipo{}, nil)
	mr := MarketRepositoryMock{}
	mr.On("FindByCodes", []string{"market-code-1"}).Return([]domain.Market{}, nil)
	mr.On("FindByIds", []domain.MarketId{}).Return([]domain.Market{}, errors.New("repository error"))

	cr := CompanyRepositoryMock{}

	ctr := CountryRepositoryMock{}
	ctr.On("FindByCodes", []string{"country-code-1"}).Return([]domain.Country{}, nil)

	sr := SectorRepositoryMock{}
	sr.On("FindByAliases", []string{"sector-alias-1"}).Return([]domain.Sector{}, nil)
	service := NewGetIposService(r, mr, cr, ctr, sr)
	query := NewGetIposQuery([]string{"market-code-1"}, []string{"country-code-1"}, []string{"sector-alias-1"}, 0)

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}


func TestService_GetIPOs_FailsWhenCompanyRepositoryFindByIdsReturnsError(t *testing.T) {
	assertion := assert.New(t)

	r := IpoRepositoryMock{}
	r.On("Find", []domain.MarketId{}, []domain.CountryId{}, []domain.SectorId{}, []domain.IndustryId{}, []domain.IpoId{}, "", uint(0), uint(20)).Return([]domain.Ipo{}, nil)

	mr := MarketRepositoryMock{}
	mr.On("FindByCodes", []string{"market-code-1"}).Return([]domain.Market{}, nil)
	mr.On("FindByIds", []domain.MarketId{}).Return([]domain.Market{}, nil)

	cr := CompanyRepositoryMock{}
	cr.On("FindByIds", []domain.CompanyId{}).Return([]domain.Company{}, errors.New("repository error"))

	ctr := CountryRepositoryMock{}
	ctr.On("FindByCodes", []string{"country-code-1"}).Return([]domain.Country{}, nil)

	sr := SectorRepositoryMock{}
	sr.On("FindByAliases", []string{"sector-alias-1"}).Return([]domain.Sector{}, nil)
	service := NewGetIposService(r, mr, cr, ctr, sr)
	query := NewGetIposQuery([]string{"market-code-1"}, []string{"country-code-1"}, []string{"sector-alias-1"}, 0)

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs_FailsWhenIpoRepositoryCountReturnsError(t *testing.T) {
	assertion := assert.New(t)

	r := IpoRepositoryMock{}
	r.On("Find", []domain.MarketId{}, []domain.CountryId{}, []domain.SectorId{}, []domain.IndustryId{}, []domain.IpoId{}, "", uint(0), uint(20)).Return([]domain.Ipo{}, nil)
	r.On("Count", []domain.MarketId{}, []domain.CountryId{}, []domain.SectorId{}, []domain.IndustryId{}, []domain.IpoId{}).Return(uint(0), errors.New("repository error"))

	mr := MarketRepositoryMock{}
	mr.On("FindByCodes", []string{"market-code-1"}).Return([]domain.Market{}, nil)
	mr.On("FindByIds", []domain.MarketId{}).Return([]domain.Market{}, nil)

	cr := CompanyRepositoryMock{}
	cr.On("FindByIds", []domain.CompanyId{}).Return([]domain.Company{}, nil)

	ctr := CountryRepositoryMock{}
	ctr.On("FindByCodes", []string{"country-code-1"}).Return([]domain.Country{}, nil)

	sr := SectorRepositoryMock{}
	sr.On("FindByAliases", []string{"sector-alias-1"}).Return([]domain.Sector{}, nil)
	service := NewGetIposService(r, mr, cr, ctr, sr)
	query := NewGetIposQuery([]string{"market-code-1"}, []string{"country-code-1"}, []string{"sector-alias-1"}, 0)

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}


func TestService_GetIPOs(t *testing.T) {
	assertion := assert.New(t)

	expectedCount := uint(6)

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

	expectedCountryReturn := make([]domain.Country, 2)
	expectedCountryReturn[0] = domain.HydrateCountry("1-country-id", "es", "Spain")
	expectedCountryReturn[1] = domain.HydrateCountry("2-country-id", "us", "USA")

	expectedSectorReturn := make([]domain.Sector, 2)
	expectedSectorReturn[0] = domain.HydrateSector("1-sector-id", "sector", "Sector")
	expectedSectorReturn[1] = domain.HydrateSector("2-sector-id", "2-sector", "Sector 2" )

	expectedMarketIdInput := make([]domain.MarketId, 2)
	expectedMarketIdInput[0] = "1-market-id"
	expectedMarketIdInput[1] = "2-market-id"

	expectedCountryIdInput := make([]domain.CountryId, 2)
	expectedCountryIdInput[0] = "1-country-id"
	expectedCountryIdInput[1] = "2-country-id"

	expectedSectorIdInput := make([]domain.SectorId, 2)
	expectedSectorIdInput[0] = "1-sector-id"
	expectedSectorIdInput[1] = "2-sector-id"

	expectedCompanyIdInput := make([]domain.CompanyId, 3)
	expectedCompanyIdInput[0] = "1-company-id"
	expectedCompanyIdInput[1] = "2-company-id"
	expectedCompanyIdInput[2] = "3-company-id"

	r := IpoRepositoryMock{}
	r.On("Find", expectedMarketIdInput, expectedCountryIdInput, expectedSectorIdInput, []domain.IndustryId{}, []domain.IpoId{}, "", uint(20), uint(20)).Return(expectedIpoReturn, nil)
	r.On("Count", expectedMarketIdInput, expectedCountryIdInput, expectedSectorIdInput, []domain.IndustryId{}, []domain.IpoId{}).Return(expectedCount, nil)

	mr := MarketRepositoryMock{}
	mr.On("FindByCodes", []string{"market-code-1", "market-code-2"}).Return(expectedMarketReturn, nil)
	mr.On("FindByIds", expectedMarketIdInput).Return(expectedMarketReturn, nil)

	cr := CompanyRepositoryMock{}
	cr.On("FindByIds", expectedCompanyIdInput).Return(expectedCompanyReturn, nil)

	ctr := CountryRepositoryMock{}
	ctr.On("FindByCodes", []string{"es", "us"}).Return(expectedCountryReturn, nil)

	sr := SectorRepositoryMock{}
	sr.On("FindByAliases", []string{"sector", "2-sector"}).Return(expectedSectorReturn, nil)

	service := NewGetIposService(r, mr, cr, ctr, sr)
	query := NewGetIposQuery([]string{"market-code-1", "market-code-2"}, []string{"es", "us"}, []string{"sector", "2-sector"}, 1)

	response, err := service.Run(query)

	assertion.NotNil(response)

	count, ipos, markets, companies := response.Get()
	assertion.Equal(expectedCount, count)
	assertion.Equal(expectedIpoReturn, ipos)
	assertion.Equal(expectedMarketReturn, markets)
	assertion.Equal(expectedCompanyReturn, companies)
	assertion.Nil(err)
}
