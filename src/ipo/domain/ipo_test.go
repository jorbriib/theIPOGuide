package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHydrateIpo(t *testing.T) {
	assertion := assert.New(t)

	company := HydrateCompany("PINS", "Pinterest")
	country := HydrateCountry("US", "USA")
	market := HydrateMarket("NQ", "Nasdaq", country)
	now := time.Now()

	ipo := HydrateIpo("NQ", company, market, &now)
	assertion.NotNil(ipo)
}

func TestIpo_Id(t *testing.T) {
	assertion := assert.New(t)

	company := HydrateCompany("PINS", "Pinterest")
	country := HydrateCountry("US", "USA")
	market := HydrateMarket("NQ", "Nasdaq", country)
	now := time.Now()

	ipo := HydrateIpo("29392-32929da", company, market, &now)

	assertion.Equal(ID("29392-32929da"), ipo.Id())
}

func TestIpo_CompanyI(t *testing.T) {
	assertion := assert.New(t)

	company := HydrateCompany("PINS", "Pinterest")
	country := HydrateCountry("US", "USA")
	market := HydrateMarket("NQ", "Nasdaq", country)
	now := time.Now()

	ipo := HydrateIpo("NQ", company, market, &now)

	assertion.Equal(company, ipo.Company())
}

func TestIpo_Market(t *testing.T) {
	assertion := assert.New(t)

	company := HydrateCompany("PINS", "Pinterest")
	country := HydrateCountry("US", "USA")
	market := HydrateMarket("NQ", "Nasdaq", country)
	now := time.Now()

	ipo := HydrateIpo("NQ", company, market, &now)

	assertion.Equal(market, ipo.Market())
}

func TestIpo_ExpectedDate(t *testing.T) {
	assertion := assert.New(t)

	company := HydrateCompany("PINS", "Pinterest")
	country := HydrateCountry("US", "USA")
	market := HydrateMarket("NQ", "Nasdaq", country)
	now := time.Now()

	ipo := HydrateIpo("NQ", company, market, &now)

	assertion.Equal(&now, ipo.ExpectedDate())
}

func TestIpo_ToString(t *testing.T) {
	assertion := assert.New(t)

	company := HydrateCompany("PINS", "Pinterest")
	country := HydrateCountry("US", "USA")
	market := HydrateMarket("NQ", "Nasdaq", country)
	now := time.Now()

	ipo := HydrateIpo("NQ", company, market, &now)

	assertion.Equal("Pinterest (PINS) in Nasdaq", ipo.ToString())
}