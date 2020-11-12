package application_test

import (
	"errors"
	. "github.com/jorbriib/theIPOGuide/backend/src/application"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGetRelatedIposService(t *testing.T) {
	mr := MarketRepositoryMock{}
	ctr := CountryRepositoryMock{}
	sr := SectorRepositoryMock{}
	service := NewGetRelatedIposService(mr, ctr, sr)
	assert.NotNil(t, service)
}

func TestNewGetRelatedIposQuery(t *testing.T) {
	query := NewGetRelatedIposQuery()
	assert.NotNil(t, query)
}

func TestGetRelatedIposService_Run_FailsWhenMarketRepositoryAllReturnsError(t *testing.T) {
	assertion := assert.New(t)

	mr := MarketRepositoryMock{}
	mr.On("All").Return([]domain.Market{}, errors.New("repository error"))

	ctr := CountryRepositoryMock{}
	sr := SectorRepositoryMock{}

	service := NewGetRelatedIposService(mr, ctr, sr)
	query := NewGetRelatedIposQuery()

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}


func TestGetRelatedIposService_Run_FailsWhenCountryRepositoryAllReturnsError(t *testing.T) {
	assertion := assert.New(t)

	mr := MarketRepositoryMock{}
	mr.On("All").Return([]domain.Market{}, nil)

	ctr := CountryRepositoryMock{}
	ctr.On("All").Return([]domain.Country{}, errors.New("repository error"))

	sr := SectorRepositoryMock{}

	service := NewGetRelatedIposService(mr, ctr, sr)
	query := NewGetRelatedIposQuery()

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}


func TestGetRelatedIposService_Run_FailsWhenSectorRepositoryAllReturnsError(t *testing.T) {
	assertion := assert.New(t)

	mr := MarketRepositoryMock{}
	mr.On("All").Return([]domain.Market{}, nil)

	ctr := CountryRepositoryMock{}
	ctr.On("All").Return([]domain.Country{}, nil)

	sr := SectorRepositoryMock{}
	sr.On("All").Return([]domain.Sector{}, errors.New("repository error"))

	service := NewGetRelatedIposService(mr, ctr, sr)
	query := NewGetRelatedIposQuery()

	response, err := service.Run(query)

	assertion.Nil(response)
	assertion.NotNil(err)
}


func TestGetRelatedIposService_Run(t *testing.T) {
	assertion := assert.New(t)

	expectedMarketReturn := make([]domain.Market, 2)
	expectedMarketReturn[0] = domain.HydrateMarket("1-market-id", "market-code-1", "Market 1", domain.HydrateCurrency("", "", ""), "image", 111)
	expectedMarketReturn[1] = domain.HydrateMarket("2-market-id", "market-code-2", "Market 2", domain.HydrateCurrency("", "", ""), "image", 222)


	expectedCountryReturn := make([]domain.Country, 2)
	expectedCountryReturn[0] = domain.HydrateCountry("1-country-id", "es", "Spain", "image", 11)
	expectedCountryReturn[1] = domain.HydrateCountry("2-country-id", "us", "USA", "image", 14)

	expectedSectorReturn := make([]domain.Sector, 2)
	expectedSectorReturn[0] = domain.HydrateSector("1-sector-id", "sector", "Sector", "image", 11)
	expectedSectorReturn[1] = domain.HydrateSector("2-sector-id", "2-sector", "Sector 2" , "image", 156)


	mr := MarketRepositoryMock{}
	mr.On("All").Return(expectedMarketReturn, nil)

	ctr := CountryRepositoryMock{}
	ctr.On("All").Return(expectedCountryReturn, nil)

	sr := SectorRepositoryMock{}
	sr.On("All").Return(expectedSectorReturn, nil)

	service := NewGetRelatedIposService(mr, ctr, sr)
	query := NewGetRelatedIposQuery()

	response, err := service.Run(query)

	assertion.Nil(err)

	assertion.NotNil(response)

	markets, countries, sectors := response.Get()
	assertion.Equal(expectedMarketReturn, markets)
	assertion.Equal(expectedCountryReturn, countries)
	assertion.Equal(expectedSectorReturn, sectors)
}