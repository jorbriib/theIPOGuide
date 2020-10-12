package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateCompany(t *testing.T) {
	assertion := assert.New(t)
	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	company := HydrateCompany(
		CompanyId("uuid"),
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
		"http://logo.com",
	)
	assertion.NotNil(company)
}


func TestCompany_Id(t *testing.T) {
	assertion := assert.New(t)
	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	company := HydrateCompany(
		CompanyId("uuid"),
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
		"http://logo.com",
	)
	assertion.Equal(CompanyId("uuid"), company.Id())
}

func TestCompany_Symbol(t *testing.T) {
	assertion := assert.New(t)
	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	company := HydrateCompany(
		CompanyId("uuid"),
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
		"http://logo.com",
	)
	assertion.Equal("PINS", company.Symbol())
}

func TestCompany_Name(t *testing.T) {
	assertion := assert.New(t)
	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	company := HydrateCompany(
		CompanyId("uuid"),
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
		"http://logo.com",
	)
	assertion.Equal("Pinterest", company.Name())
}

func TestCompany_Sector(t *testing.T) {
	assertion := assert.New(t)
	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	company := HydrateCompany(
		CompanyId("uuid"),
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
		"http://logo.com",
	)
	assertion.Equal(sector, company.Sector())
}

func TestCompany_Country(t *testing.T) {
	assertion := assert.New(t)
	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	company := HydrateCompany(
		CompanyId("uuid"),
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
		"http://logo.com",
	)
	assertion.Equal(country, company.Country())
}

func TestCompany_Address(t *testing.T) {
	assertion := assert.New(t)
	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	company := HydrateCompany(
		CompanyId("uuid"),
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
		"http://logo.com",
	)
	assertion.Equal("Av.2", company.Address())
}

func TestCompany_LogoUrl(t *testing.T) {
	assertion := assert.New(t)
	sector := HydrateSector("Communication")
	country := HydrateCountry("US", "USA")
	company := HydrateCompany(
		CompanyId("uuid"),
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
		"http://logo.com",
	)
	assertion.Equal("http://logo.com", company.LogoUrl())
}
