package application_test

import (
	"errors"
	. "github.com/jorbriib/theIPOGuide/backend/src/application"
	. "github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sort"
	"testing"
)

func TestNewGetIpoService(t *testing.T) {
	r := IpoRepositoryMock{}
	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}
	service := NewGetIpoService(r, mr, cr)
	assert.NotNil(t, service)
}

func TestNewGetIpoQuery(t *testing.T) {
	query := NewGetIpoQuery("alias")
	assert.Equal(t, "alias", query.Alias())
}

func TestService_GetIPO_FailsWhenMarketRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)

	alias := "pinterest"
	expectedReturn := HydrateIpo("1-ipo-id", "1-alias", "intro", "1-market-id", "1-company-id", 0, 0, 0, nil)

	r := IpoRepositoryMock{}
	r.On("GetByAlias", alias).Return(&expectedReturn, nil)

	expectedMarketIdInput := MarketId("1-market-id")

	mr := MarketRepositoryMock{}
	mr.On("GetById", expectedMarketIdInput).Return(&Market{}, errors.New("repository error"))

	cr := CompanyRepositoryMock{}

	service := NewGetIpoService(r, mr, cr)
	query := NewGetIpoQuery(alias)
	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPO_FailsWhenCompanyRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)
	alias := "pinterest"

	expectedIpoReturn := HydrateIpo("1-ipo-id", "1-alias", "intro", "1-market-id", "1-company-id", 0, 0, 0, nil)
	expectedMarketReturn := HydrateMarket("1-market-id", "", "", HydrateCurrency("", "", ""))

	r := IpoRepositoryMock{}
	r.On("GetByAlias", alias).Return(&expectedIpoReturn, nil)

	expectedMarketIdInput := MarketId("1-market-id")

	mr := MarketRepositoryMock{}
	mr.On("GetById", expectedMarketIdInput).Return(&expectedMarketReturn, nil)

	expectedCompanyIdInput := CompanyId("1-company-id")

	cr := CompanyRepositoryMock{}
	cr.On("GetById", expectedCompanyIdInput).Return(&Company{}, errors.New("repository error"))

	service := NewGetIpoService(r, mr, cr)
	query := NewGetIpoQuery(alias)
	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPO_ReturnsNilWhenCompanyRepositoryReturnsNil(t *testing.T) {

	assertion := assert.New(t)
	alias := "pinterest"

	r := IpoRepositoryMock{}
	r.On("GetByAlias", alias).Return(nil, nil).Once()

	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}

	service := NewGetIpoService(r, mr, cr)
	query := NewGetIpoQuery(alias)
	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.Nil(err)
}

func TestService_GetIPO(t *testing.T) {
	assertion := assert.New(t)
	alias := "pinterest"

	expectedIpoReturn := HydrateIpo("1-ipo-id", "1-alias", "intro", "1-market-id", "1-company-id", 0, 0, 0, nil)
	expectedMarketReturn := HydrateMarket("1-market-id", "", "", HydrateCurrency("", "", ""))

	expectedCompanyReturn := HydrateCompany(
		"1-company-id", "", "",
		HydrateSector("1-sector", "1-alias", "sector"),
		HydrateIndustry("1-industry", "1-alias", "industry"), "",
		HydrateCountry("1-country", "1-code", "country"), "",
		"", "", 0, "", "",
		"", "", "", "", 2000, "", "", "", "", "")

	r := IpoRepositoryMock{}
	r.On("GetByAlias", alias).Return(&expectedIpoReturn, nil)

	expectedMarketIdInput := MarketId("1-market-id")
	mr := MarketRepositoryMock{}
	mr.On("GetById", expectedMarketIdInput).Return(&expectedMarketReturn, nil)

	expectedCompanyIdInput := CompanyId("1-company-id")
	cr := CompanyRepositoryMock{}
	cr.On("GetById", expectedCompanyIdInput).Return(&expectedCompanyReturn, nil)

	service := NewGetIpoService(r, mr, cr)
	query := NewGetIpoQuery(alias)
	response, err := service.Run(query)

	assertion.NotNil(response)

	ipo, market, company := response.Get()
	assertion.Equal(&expectedIpoReturn, ipo)
	assertion.Equal(&expectedMarketReturn, market)
	assertion.Equal(&expectedCompanyReturn, company)
	assertion.Nil(err)
}


type IpoRepositoryMock struct {
	mock.Mock
}

func (r IpoRepositoryMock) Find(marketIds []MarketId, countryIds []CountryId, sectorIds []SectorId, industryIds []IndustryId, blackList []IpoId, sortBy string, offset uint, limit uint) ([]Ipo, error) {
	args := r.Called(marketIds, countryIds, sectorIds, industryIds, blackList, sortBy, offset, limit)
	return args.Get(0).([]Ipo), args.Error(1)
}

func (r IpoRepositoryMock) GetByAlias(alias string) (*Ipo, error) {
	args := r.Called(alias)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Ipo), args.Error(1)
}

func (r IpoRepositoryMock) Count(marketIds []MarketId, countryIds []CountryId, sectorIds []SectorId, industryIds []IndustryId, blackList []IpoId) (uint, error) {
	args := r.Called(marketIds, countryIds, sectorIds, industryIds, blackList)
	return args.Get(0).(uint), args.Error(1)
}

func (r IpoRepositoryMock) SearchByText(text string) ([]Ipo, error) {
	args := r.Called(text)
	return args.Get(0).([]Ipo), args.Error(1)
}

type MarketRepositoryMock struct {
	mock.Mock
}

func (r MarketRepositoryMock) FindByIds(ids []MarketId) ([]Market, error) {
	sort.SliceStable(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	args := r.Called(ids)
	return args.Get(0).([]Market), args.Error(1)
}

func (r MarketRepositoryMock) GetById(id MarketId) (*Market, error) {
	args := r.Called(id)
	return args.Get(0).(*Market), args.Error(1)
}

func (r MarketRepositoryMock) All() ([]Market, error) {
	args := r.Called()
	return args.Get(0).([]Market), args.Error(1)
}

func (r MarketRepositoryMock) FindByCodes(codes []string) ([]Market, error) {
	args := r.Called(codes)
	return args.Get(0).([]Market), args.Error(1)
}

type CompanyRepositoryMock struct {
	mock.Mock
}

func (r CompanyRepositoryMock) FindByIds(ids []CompanyId) ([]Company, error) {
	sort.SliceStable(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	args := r.Called(ids)
	return args.Get(0).([]Company), args.Error(1)
}

func (r CompanyRepositoryMock) GetById(id CompanyId) (*Company, error) {
	args := r.Called(id)
	return args.Get(0).(*Company), args.Error(1)
}

type CountryRepositoryMock struct {
	mock.Mock
}

func (c CountryRepositoryMock) All() ([]Country, error) {
	args := c.Called()
	return args.Get(0).([]Country), args.Error(1)
}

func (c CountryRepositoryMock) FindByCodes(codes []string) ([]Country, error) {
	args := c.Called(codes)
	return args.Get(0).([]Country), args.Error(1)
}

type SectorRepositoryMock struct {
	mock.Mock
}

func (s SectorRepositoryMock) All() ([]Sector, error) {
	args := s.Called()
	return args.Get(0).([]Sector), args.Error(1)
}

func (s SectorRepositoryMock) FindByAliases(aliases []string) ([]Sector, error) {
	args := s.Called(aliases)
	return args.Get(0).([]Sector), args.Error(1)
}
