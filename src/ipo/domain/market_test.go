package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateMarket(t *testing.T) {
	assertion := assert.New(t)
	country := HydrateCountry("US", "USA")
	market := HydrateMarket("NQ", "Nasdaq", country)
	assertion.NotNil(market)
}

func TestMarket_Symbol(t *testing.T) {
	assertion := assert.New(t)
	country := HydrateCountry("US", "USA")
	market := HydrateMarket("NQ", "Nasdaq", country)

	assertion.Equal("NQ", market.Code())
}

func TestMarket_Name(t *testing.T) {
	assertion := assert.New(t)
	country := HydrateCountry("US", "USA")
	market := HydrateMarket("NQ", "Nasdaq", country)

	assertion.Equal("Nasdaq", market.Name())
}


func TestMarket_Country(t *testing.T) {
	assertion := assert.New(t)
	country := HydrateCountry("US", "USA")
	market := HydrateMarket("NQ", "Nasdaq", country)

	assertion.Equal(country, market.Country())
}