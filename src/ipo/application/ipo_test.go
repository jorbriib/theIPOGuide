package application_test

import (
	"errors"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sort"
	"testing"
)

func TestNewService(t *testing.T) {
	assertion := assert.New(t)
	r := IpoRepositoryMock{}
	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}
	service := application.NewService(r, mr, cr)
	assertion.NotNil(service)
}

func TestService_GetIPOs_FailsWhenIpoRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)

	r := IpoRepositoryMock{}
		r.On("Find").Return([]domain.Ipo{}, errors.New("repository error"))

	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}

	service := application.NewService(r, mr, cr)
	query := application.NewGetIposQuery()
	response, err := service.GetIPOs(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs_FailsWhenMarketRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)

	expectedIpoReturn := make([]domain.Ipo, 2)
	expectedIpoReturn[0] = domain.HydrateIpo("1-ipo-id", "1-alias", "1-market-id", "1-company-id", 0, 0, 0, nil)
	expectedIpoReturn[1] = domain.HydrateIpo("2-ipo-id", "2-alias", "2-market-id", "2-company-id", 0, 0, 0, nil)


	r := IpoRepositoryMock{}
	r.On("Find").Return(expectedIpoReturn, nil)

	expectedMarketIdInput := make([]domain.MarketId, 2)
	expectedMarketIdInput[0] = "1-market-id"
	expectedMarketIdInput[1] = "2-market-id"

	mr := MarketRepositoryMock{}
	mr.On("FindByIds", expectedMarketIdInput).Return([]domain.Market{}, errors.New("repository error"))

	cr := CompanyRepositoryMock{}

	service := application.NewService(r, mr, cr)
	query := application.NewGetIposQuery()
	response, err := service.GetIPOs(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs_FailsWhenCompanyRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)

	expectedIpoReturn := make([]domain.Ipo, 3)
	expectedIpoReturn[0] = domain.HydrateIpo("1-ipo-id", "1-alias", "1-market-id", "1-company-id", 0, 0, 0, nil)
	expectedIpoReturn[1] = domain.HydrateIpo("2-ipo-id", "2-alias", "2-market-id", "2-company-id", 0, 0, 0, nil)
	expectedIpoReturn[2] = domain.HydrateIpo("3-ipo-id", "3-alias", "2-market-id", "3-company-id", 0, 0, 0, nil)

	expectedMarketReturn := make([]domain.Market, 2)
	expectedMarketReturn[0] = domain.HydrateMarket("1-market-id", "", "", domain.HydrateCurrency("", "", ""))
	expectedMarketReturn[1] = domain.HydrateMarket("2-market-id", "", "", domain.HydrateCurrency("", "", ""))


	r := IpoRepositoryMock{}
	r.On("Find").Return(expectedIpoReturn, nil)

	expectedMarketIdInput := make([]domain.MarketId, 2)
	expectedMarketIdInput[0] = "1-market-id"
	expectedMarketIdInput[1] = "2-market-id"

	mr := MarketRepositoryMock{}
	mr.On("FindByIds", expectedMarketIdInput).Return(expectedMarketReturn, nil)

	expectedCompanyIdInput := make([]domain.CompanyId, 3)
	expectedCompanyIdInput[0] = domain.CompanyId("1-company-id")
	expectedCompanyIdInput[1] = domain.CompanyId("2-company-id")
	expectedCompanyIdInput[2] = domain.CompanyId("3-company-id")

	cr := CompanyRepositoryMock{}
	cr.On("FindByIds", expectedCompanyIdInput).Return([]domain.Company{}, errors.New("repository error"))

	service := application.NewService(r, mr, cr)
	query := application.NewGetIposQuery()
	response, err := service.GetIPOs(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs(t *testing.T) {
	assertion := assert.New(t)

	expectedIpoReturn := make([]domain.Ipo, 3)
	expectedIpoReturn[0] = domain.HydrateIpo("1-ipo-id", "1-alias", "1-market-id", "1-company-id", 0, 0, 0, nil)
	expectedIpoReturn[1] = domain.HydrateIpo("2-ipo-id", "2-alias", "2-market-id", "2-company-id", 0, 0, 0, nil)
	expectedIpoReturn[2] = domain.HydrateIpo("3-ipo-id", "3-alias", "2-market-id", "3-company-id", 0, 0, 0, nil)

	expectedMarketReturn := make([]domain.Market, 2)
	expectedMarketReturn[0] = domain.HydrateMarket("1-market-id", "", "", domain.HydrateCurrency("", "", ""))
	expectedMarketReturn[1] = domain.HydrateMarket("2-market-id", "", "", domain.HydrateCurrency("", "", ""))

	expectedCompanyReturn := make([]domain.Company, 3)
	expectedCompanyReturn[0] = domain.HydrateCompany(
		"1-company-id", "", "",
		domain.HydrateSector(""), domain.HydrateIndustry(""), "", domain.HydrateCountry("", ""), "",
		"", "", 0, "", "",
		"", "", "", "", 2000, "", "", "", "", "")
	expectedCompanyReturn[1] = domain.HydrateCompany(
		"2-company-id", "", "",
		domain.HydrateSector(""), domain.HydrateIndustry(""), "", domain.HydrateCountry("", ""), "",
		"", "", 0, "", "",
		"", "", "", "", 2000, "", "", "", "", "")
	expectedCompanyReturn[2] = domain.HydrateCompany(
		"3-company-id", "", "",
		domain.HydrateSector(""), domain.HydrateIndustry(""), "", domain.HydrateCountry("", ""), "",
		"", "", 0, "", "",
		"", "", "", "", 2000, "", "", "", "", "")

	r := IpoRepositoryMock{}
	r.On("Find").Return(expectedIpoReturn, nil)

	expectedMarketIdInput := make([]domain.MarketId, 2)
	expectedMarketIdInput[0] = "1-market-id"
	expectedMarketIdInput[1] = "2-market-id"

	mr := MarketRepositoryMock{}
	mr.On("FindByIds", expectedMarketIdInput).Return(expectedMarketReturn, nil)

	expectedCompanyIdInput := make([]domain.CompanyId, 3)
	expectedCompanyIdInput[0] = "1-company-id"
	expectedCompanyIdInput[1] = "2-company-id"
	expectedCompanyIdInput[2] = "3-company-id"

	cr := CompanyRepositoryMock{}
	cr.On("FindByIds", expectedCompanyIdInput).Return(expectedCompanyReturn, nil)

	service := application.NewService(r, mr, cr)
	query := application.NewGetIposQuery()
	response, err := service.GetIPOs(query)

	assertion.NotNil(response)

	ipos, markets, companies := response.Get()
	assertion.Equal(expectedIpoReturn, ipos)
	assertion.Equal(expectedMarketReturn, markets)
	assertion.Equal(expectedCompanyReturn, companies)
	assertion.Nil(err)
}

func TestService_GetIPO_FailsWhenIpoRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)

	alias := "pinterest"

	r := IpoRepositoryMock{}
	r.On("GetByAlias", alias).Return(&domain.Ipo{}, errors.New("repository error"))

	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}

	service := application.NewService(r, mr, cr)
	query := application.NewGetIpoQuery(alias)
	response, err := service.GetIPO(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPO_FailsWhenMarketRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)

	alias := "pinterest"
	expectedReturn := domain.HydrateIpo("1-ipo-id", "1-alias", "1-market-id", "1-company-id", 0, 0, 0, nil)

	r := IpoRepositoryMock{}
	r.On("GetByAlias", alias).Return(&expectedReturn, nil)

	expectedMarketIdInput := domain.MarketId("1-market-id")

	mr := MarketRepositoryMock{}
	mr.On("GetById", expectedMarketIdInput).Return(&domain.Market{}, errors.New("repository error"))

	cr := CompanyRepositoryMock{}

	service := application.NewService(r, mr, cr)
	query := application.NewGetIpoQuery(alias)
	response, err := service.GetIPO(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPO_FailsWhenCompanyRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)
	alias := "pinterest"

	expectedIpoReturn := domain.HydrateIpo("1-ipo-id", "1-alias", "1-market-id", "1-company-id", 0, 0, 0, nil)
	expectedMarketReturn := domain.HydrateMarket("1-market-id", "", "", domain.HydrateCurrency("", "", ""))

	r := IpoRepositoryMock{}
	r.On("GetByAlias", alias).Return(&expectedIpoReturn, nil)

	expectedMarketIdInput := domain.MarketId("1-market-id")

	mr := MarketRepositoryMock{}
	mr.On("GetById", expectedMarketIdInput).Return(&expectedMarketReturn, nil)

	expectedCompanyIdInput := domain.CompanyId("1-company-id")

	cr := CompanyRepositoryMock{}
	cr.On("GetById", expectedCompanyIdInput).Return(&domain.Company{}, errors.New("repository error"))

	service := application.NewService(r, mr, cr)
	query := application.NewGetIpoQuery(alias)
	response, err := service.GetIPO(query)

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

	service := application.NewService(r, mr, cr)
	query := application.NewGetIpoQuery(alias)
	response, err := service.GetIPO(query)

	assertion.Nil(response)
	assertion.Nil(err)
}


func TestService_GetIPO(t *testing.T) {
	assertion := assert.New(t)
	alias := "pinterest"

	expectedIpoReturn := domain.HydrateIpo("1-ipo-id", "1-alias", "1-market-id", "1-company-id", 0, 0, 0, nil)
	expectedMarketReturn := domain.HydrateMarket("1-market-id", "", "", domain.HydrateCurrency("", "", ""))

	expectedCompanyReturn := domain.HydrateCompany(
		"1-company-id", "", "",
		domain.HydrateSector(""), domain.HydrateIndustry(""), "", domain.HydrateCountry("", ""), "",
		"", "", 0, "", "",
		"", "", "", "", 2000, "", "", "", "", "")

	r := IpoRepositoryMock{}
	r.On("GetByAlias", alias).Return(&expectedIpoReturn, nil)

	expectedMarketIdInput := domain.MarketId("1-market-id")
	mr := MarketRepositoryMock{}
	mr.On("GetById", expectedMarketIdInput).Return(&expectedMarketReturn, nil)

	expectedCompanyIdInput := domain.CompanyId("1-company-id")
	cr := CompanyRepositoryMock{}
	cr.On("GetById", expectedCompanyIdInput).Return(&expectedCompanyReturn, nil)

	service := application.NewService(r, mr, cr)
	query := application.NewGetIpoQuery(alias)
	response, err := service.GetIPO(query)

	assertion.NotNil(response)

	ipo, market, company := response.Get()
	assertion.Equal(&expectedIpoReturn, ipo)
	assertion.Equal(&expectedMarketReturn, market)
	assertion.Equal(&expectedCompanyReturn, company)
	assertion.Nil(err)
}

func TestNewGetIposQuery(t *testing.T) {
	assertion := assert.New(t)
	query := application.NewGetIposQuery()
	assertion.NotNil(query)
}

func TestNewGetIpoQuery(t *testing.T) {
	assertion := assert.New(t)
	alias := "pinterest"

	query := application.NewGetIpoQuery(alias)
	assertion.NotNil(query)
	assertion.Equal(alias, query.Alias())
}

type IpoRepositoryMock struct {
	mock.Mock
}

func (r IpoRepositoryMock) Find() ([]domain.Ipo, error) {
	args := r.Called()
	return args.Get(0).([]domain.Ipo), args.Error(1)
}

func (r IpoRepositoryMock) GetByAlias(alias string) (*domain.Ipo, error) {
	args := r.Called(alias)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Ipo), args.Error(1)
}

type MarketRepositoryMock struct {
	mock.Mock
}

func (r MarketRepositoryMock) FindByIds(ids []domain.MarketId) ([]domain.Market, error) {
	sort.SliceStable(ids, func(i, j int) bool{
		return ids[i] < ids[j]
	})
	args := r.Called(ids)
	return args.Get(0).([]domain.Market), args.Error(1)
}

func (r MarketRepositoryMock) GetById(id domain.MarketId) (*domain.Market, error) {
	args := r.Called(id)
	return args.Get(0).(*domain.Market), args.Error(1)
}

type CompanyRepositoryMock struct {
	mock.Mock
}

func (r CompanyRepositoryMock) FindByIds(ids []domain.CompanyId) ([]domain.Company, error) {
	sort.SliceStable(ids, func(i, j int) bool{
		return ids[i] < ids[j]
	})
	args := r.Called(ids)
	return args.Get(0).([]domain.Company), args.Error(1)
}

func (r CompanyRepositoryMock) GetById(id domain.CompanyId) (*domain.Company, error) {
	args := r.Called(id)
	return args.Get(0).(*domain.Company), args.Error(1)
}
