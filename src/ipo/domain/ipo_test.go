package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHydrateIpo(t *testing.T) {
	assertion := assert.New(t)

	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	media := HydrateMedia("image", "jpg", "http://media.com", "logo")
	company := HydrateCompany(
		"PINS",
		"Pinterest",
		sector,
		"Av.2",
		country,
		"93 38489390",
		"email@email.com",
		"https://website.com",
		10000,
		"description",
		2002,
		"Tomas Cook",
		"March 31",
		"http://ipourl.com",
		"http://commission.com",
		media,
	)

	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)
	now := time.Now()

	ipo := HydrateIpo("NQ", market, company, 3222, 3444, 10029039, &now)
	assertion.NotNil(ipo)
}

func TestIpo_Id(t *testing.T) {
	assertion := assert.New(t)

	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	media := HydrateMedia("image", "jpg", "http://media.com", "logo")
	company := HydrateCompany(
		"PINS",
		"Pinterest",
		sector,
		"Av.2",
		country,
		"93 38489390",
		"email@email.com",
		"https://website.com",
		10000,
		"description",
		2002,
		"Tomas Cook",
		"March 31",
		"http://ipourl.com",
		"http://commission.com",
		media,
	)

	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)
	now := time.Now()

	ipo := HydrateIpo("NQ", market, company, 3222, 3444, 10029039, &now)

	assertion.Equal(ID("29392-32929da"), ipo.Id())
}

func TestIpo_Company(t *testing.T) {
	assertion := assert.New(t)

	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	media := HydrateMedia("image", "jpg", "http://media.com", "logo")
	company := HydrateCompany(
		"PINS",
		"Pinterest",
		sector,
		"Av.2",
		country,
		"93 38489390",
		"email@email.com",
		"https://website.com",
		10000,
		"description",
		2002,
		"Tomas Cook",
		"March 31",
		"http://ipourl.com",
		"http://commission.com",
		media,
	)

	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)
	now := time.Now()

	ipo := HydrateIpo("NQ", market, company, 3222, 3444, 10029039, &now)

	assertion.Equal(company, ipo.Company())
}

func TestIpo_Market(t *testing.T) {
	assertion := assert.New(t)

	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	media := HydrateMedia("image", "jpg", "http://media.com", "logo")
	company := HydrateCompany(
		"PINS",
		"Pinterest",
		sector,
		"Av.2",
		country,
		"93 38489390",
		"email@email.com",
		"https://website.com",
		10000,
		"description",
		2002,
		"Tomas Cook",
		"March 31",
		"http://ipourl.com",
		"http://commission.com",
		media,
	)

	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)
	now := time.Now()

	ipo := HydrateIpo("NQ", market, company, 3222, 3444, 10029039, &now)

	assertion.Equal(market, ipo.Market())
}

func TestIpo_PriceCentsFrom(t *testing.T) {
	assertion := assert.New(t)

	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	media := HydrateMedia("image", "jpg", "http://media.com", "logo")
	company := HydrateCompany(
		"PINS",
		"Pinterest",
		sector,
		"Av.2",
		country,
		"93 38489390",
		"email@email.com",
		"https://website.com",
		10000,
		"description",
		2002,
		"Tomas Cook",
		"March 31",
		"http://ipourl.com",
		"http://commission.com",
		media,
	)

	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)
	now := time.Now()

	ipo := HydrateIpo("NQ", market, company, 3222, 3444, 10029039, &now)

	assertion.Equal(3222, ipo.PriceCentsFrom())
}

func TestIpo_PriceCentsTo(t *testing.T) {
	assertion := assert.New(t)

	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	media := HydrateMedia("image", "jpg", "http://media.com", "logo")
	company := HydrateCompany(
		"PINS",
		"Pinterest",
		sector,
		"Av.2",
		country,
		"93 38489390",
		"email@email.com",
		"https://website.com",
		10000,
		"description",
		2002,
		"Tomas Cook",
		"March 31",
		"http://ipourl.com",
		"http://commission.com",
		media,
	)

	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)
	now := time.Now()

	ipo := HydrateIpo("NQ", market, company, 3222, 3444, 10029039, &now)

	assertion.Equal(3444, ipo.PriceCentsTo())
}

func TestIpo_Shares(t *testing.T) {
	assertion := assert.New(t)

	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	media := HydrateMedia("image", "jpg", "http://media.com", "logo")
	company := HydrateCompany(
		"PINS",
		"Pinterest",
		sector,
		"Av.2",
		country,
		"93 38489390",
		"email@email.com",
		"https://website.com",
		10000,
		"description",
		2002,
		"Tomas Cook",
		"March 31",
		"http://ipourl.com",
		"http://commission.com",
		media,
	)

	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)
	now := time.Now()

	ipo := HydrateIpo("NQ", market, company, 3222, 3444, 10029039, &now)

	assertion.Equal(10029039, ipo.Shares())
}


func TestIpo_ExpectedDate(t *testing.T) {
	assertion := assert.New(t)

	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	media := HydrateMedia("image", "jpg", "http://media.com", "logo")
	company := HydrateCompany(
		"PINS",
		"Pinterest",
		sector,
		"Av.2",
		country,
		"93 38489390",
		"email@email.com",
		"https://website.com",
		10000,
		"description",
		2002,
		"Tomas Cook",
		"March 31",
		"http://ipourl.com",
		"http://commission.com",
		media,
	)

	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)
	now := time.Now()

	ipo := HydrateIpo("NQ", market, company, 3222, 3444, 10029039, &now)

	assertion.Equal(&now, ipo.ExpectedDate())
}

func TestIpo_ToString(t *testing.T) {
	assertion := assert.New(t)

	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	media := HydrateMedia("image", "jpg", "http://media.com", "logo")
	company := HydrateCompany(
		"PINS",
		"Pinterest",
		sector,
		"Av.2",
		country,
		"93 38489390",
		"email@email.com",
		"https://website.com",
		10000,
		"description",
		2002,
		"Tomas Cook",
		"March 31",
		"http://ipourl.com",
		"http://commission.com",
		media,
	)

	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)
	now := time.Now()

	ipo := HydrateIpo("NQ", market, company, 3222, 3444, 10029039, &now)

	assertion.Equal("Pinterest (PINS) in Nasdaq", ipo.ToString())
}