package domain_test

import (
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateCompany(t *testing.T) {
	assertion := assert.New(t)
	sector := domain.HydrateSector("Communication")
	industry := domain.HydrateIndustry("Tech")
	country := domain.HydrateCountry("US", "USA")
	company := domain.HydrateCompany(
		"uuid",
		"PINS",
		"Pinterest",
		sector,
		industry,
		"Av.2",
		country,
		"93 38489390",
		"email@email.com",
		"https://website.com",
		10000,
		"description",
		"http://facebook.com",
		"http://twitter.com",
		"http://linkedin.com",
		"http://pinterest.com",
		"http://instagram.com",
		2020,
		"Tomas Cook",
		"March 31",
		"http://ipourl.com",
		"http://commission.com",
		"http://logo.com",
	)
	assertion.NotNil(company)

	assertion.Equal(domain.CompanyId("uuid"), company.Id())
	assertion.Equal("PINS", company.Symbol())
	assertion.Equal("Pinterest", company.Name())
	assertion.Equal(sector, company.Sector())
	assertion.Equal(industry, company.Industry())
	assertion.Equal("Av.2", company.Address())
	assertion.Equal(country, company.Country())
	assertion.Equal("93 38489390", company.Phone())
	assertion.Equal("email@email.com", company.Email())
	assertion.Equal("https://website.com", company.Website())
	assertion.Equal(uint32(10000), company.Employees())
	assertion.Equal("description", company.Description())
	assertion.Equal("http://facebook.com", company.Facebook())
	assertion.Equal("http://twitter.com", company.Twitter())
	assertion.Equal("http://linkedin.com", company.Linkedin())
	assertion.Equal("http://pinterest.com", company.Pinterest())
	assertion.Equal("http://instagram.com", company.Instagram())
	assertion.Equal(uint16(2020), company.Founded())
	assertion.Equal("Tomas Cook", company.Ceo())
	assertion.Equal("March 31", company.FiscalYearEnd())
	assertion.Equal("http://ipourl.com", company.IpoUrl())
	assertion.Equal("http://commission.com", company.ExchangeCommissionUrl())
	assertion.Equal("http://logo.com", company.LogoUrl())
}
