package application

import (
	"errors"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewService(t *testing.T) {
	assertion := assert.New(t)
	r := IpoRepositoryMock{}
	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}
	service := NewService(r, mr, cr)
	assertion.NotNil(service)
}

func TestService_GetIPOs_FailsWhenIpoRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)

	r := IpoRepositoryMock{}
	r.On("Find").Return([]domain.Ipo{}, errors.New("repository error"))

	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}

	service := NewService(r, mr, cr)
	query := NewGetIposQuery()
	response, err := service.GetIPOs(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs_FailsWhenMarketRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)

	expectedReturn := []domain.Ipo{
		domain.HydrateIpo("1-ipo-id", "1-market-id", "1-company-id", 0, 0, 0, nil),
		domain.HydrateIpo("2-ipo-id", "2-market-id", "2-company-id", 0, 0, 0, nil),
	}

	r := IpoRepositoryMock{}
	r.On("Find").Return(expectedReturn, nil)

	expectedMarketIdInput := []domain.MarketId{
		domain.MarketId("1-market-id"),
		domain.MarketId("2-market-id"),
	}
	mr := MarketRepositoryMock{}
	mr.On("FindByIds", expectedMarketIdInput).Return([]domain.Market{}, errors.New("repository error"))

	cr := CompanyRepositoryMock{}

	service := NewService(r, mr, cr)
	query := NewGetIposQuery()
	response, err := service.GetIPOs(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs_FailsWhenCompanyRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)

	expectedIpoReturn := []domain.Ipo{
		domain.HydrateIpo("1-ipo-id", "1-market-id", "1-company-id", 0, 0, 0, nil),
		domain.HydrateIpo("2-ipo-id", "2-market-id", "2-company-id", 0, 0, 0, nil),
		domain.HydrateIpo("3-ipo-id", "2-market-id", "3-company-id", 0, 0, 0, nil),
	}

	expectedMarketReturn := []domain.Market{
		domain.HydrateMarket("1-market-id", "", "", domain.HydrateCurrency("", "", "")),
		domain.HydrateMarket("2-market-id", "", "", domain.HydrateCurrency("", "", "")),
	}

	r := IpoRepositoryMock{}
	r.On("Find").Return(expectedIpoReturn, nil)

	expectedMarketIdInput := []domain.MarketId{
		domain.MarketId("1-market-id"),
		domain.MarketId("2-market-id"),
	}
	mr := MarketRepositoryMock{}
	mr.On("FindByIds", expectedMarketIdInput).Return(expectedMarketReturn, nil)

	expectedCompanyIdInput := []domain.CompanyId{
		domain.CompanyId("1-company-id"),
		domain.CompanyId("2-company-id"),
		domain.CompanyId("3-company-id"),
	}
	cr := CompanyRepositoryMock{}
	cr.On("FindByIds", expectedCompanyIdInput).Return([]domain.Company{}, errors.New("repository error"))

	service := NewService(r, mr, cr)
	query := NewGetIposQuery()
	response, err := service.GetIPOs(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}

func TestService_GetIPOs(t *testing.T) {
	assertion := assert.New(t)

	expectedIpoReturn := []domain.Ipo{
		domain.HydrateIpo("1-ipo-id", "1-market-id", "1-company-id", 0, 0, 0, nil),
		domain.HydrateIpo("2-ipo-id", "2-market-id", "2-company-id", 0, 0, 0, nil),
		domain.HydrateIpo("3-ipo-id", "2-market-id", "3-company-id", 0, 0, 0, nil),
	}

	expectedMarketReturn := []domain.Market{
		domain.HydrateMarket("1-market-id", "", "", domain.HydrateCurrency("", "", "")),
		domain.HydrateMarket("2-market-id", "", "", domain.HydrateCurrency("", "", "")),
	}

	expectedCompanyReturn := []domain.Company{
		domain.HydrateCompany(
			"1-company-id", "", "",
			domain.HydrateSector(""), "", domain.HydrateCountry("", ""), "",
			"", "", 0, "", 2000,
			"", "", "", "", ""),
		domain.HydrateCompany(
			"2-company-id", "", "",
			domain.HydrateSector(""), "", domain.HydrateCountry("", ""), "",
			"", "", 0, "", 2000,
			"", "", "", "", ""),
		domain.HydrateCompany(
			"3-company-id", "", "",
			domain.HydrateSector(""), "", domain.HydrateCountry("", ""), "",
			"", "", 0, "", 2000,
			"", "", "", "", ""),
	}

	r := IpoRepositoryMock{}
	r.On("Find").Return(expectedIpoReturn, nil)

	expectedMarketIdInput := []domain.MarketId{
		domain.MarketId("1-market-id"),
		domain.MarketId("2-market-id"),
	}
	mr := MarketRepositoryMock{}
	mr.On("FindByIds", expectedMarketIdInput).Return(expectedMarketReturn, nil)

	expectedCompanyIdInput := []domain.CompanyId{
		domain.CompanyId("1-company-id"),
		domain.CompanyId("2-company-id"),
		domain.CompanyId("3-company-id"),
	}
	cr := CompanyRepositoryMock{}
	cr.On("FindByIds", expectedCompanyIdInput).Return(expectedCompanyReturn, nil)

	service := NewService(r, mr, cr)
	query := NewGetIposQuery()
	response, err := service.GetIPOs(query)

	assertion.NotNil(response)

	ipos, markets, companies := response.Get()
	assertion.Equal(expectedIpoReturn, ipos)
	assertion.Equal(expectedMarketReturn, markets)
	assertion.Equal(expectedCompanyReturn, companies)
	assertion.Nil(err)
}

func TestNewGetIposQuery(t *testing.T) {
	assertion := assert.New(t)
	query := NewGetIposQuery()
	assertion.NotNil(query)
}

type IpoRepositoryMock struct {
	mock.Mock
}

func (r IpoRepositoryMock) Find() ([]domain.Ipo, error) {
	args := r.Called()
	return args.Get(0).([]domain.Ipo), args.Error(1)
}

type MarketRepositoryMock struct {
	mock.Mock
}

func (r MarketRepositoryMock) FindByIds(ids []domain.MarketId) ([]domain.Market, error) {
	args := r.Called(ids)
	return args.Get(0).([]domain.Market), args.Error(1)
}

type CompanyRepositoryMock struct {
	mock.Mock
}

func (r CompanyRepositoryMock) FindByIds(ids []domain.CompanyId) ([]domain.Company, error) {
	args := r.Called(ids)
	return args.Get(0).([]domain.Company), args.Error(1)
}
