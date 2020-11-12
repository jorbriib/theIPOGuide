package domain_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateCompany(t *testing.T) {
	sector := HydrateSector("1-sector", "communication", "Communication", "image", 10)
	industry := HydrateIndustry("1-industry", "tech", "Tech")
	country := HydrateCountry("1-contry", "US", "USA", "image", 2)
	company := HydrateCompany(
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
	assert.NotNil(t, company)

	assert.Equal(t, CompanyId("uuid"), company.Id())
	assert.Equal(t, "PINS", company.Symbol())
	assert.Equal(t, "Pinterest", company.Name())
	assert.Equal(t, sector, company.Sector())
	assert.Equal(t, industry, company.Industry())
	assert.Equal(t, "Av.2", company.Address())
	assert.Equal(t, country, company.Country())
	assert.Equal(t, "93 38489390", company.Phone())
	assert.Equal(t, "email@email.com", company.Email())
	assert.Equal(t, "https://website.com", company.Website())
	assert.Equal(t, uint32(10000), company.Employees())
	assert.Equal(t, "description", company.Description())
	assert.Equal(t, "http://facebook.com", company.Facebook())
	assert.Equal(t, "http://twitter.com", company.Twitter())
	assert.Equal(t, "http://linkedin.com", company.Linkedin())
	assert.Equal(t, "http://pinterest.com", company.Pinterest())
	assert.Equal(t, "http://instagram.com", company.Instagram())
	assert.Equal(t, uint16(2020), company.Founded())
	assert.Equal(t, "Tomas Cook", company.Ceo())
	assert.Equal(t, "March 31", company.FiscalYearEnd())
	assert.Equal(t, "http://ipourl.com", company.IpoUrl())
	assert.Equal(t, "http://commission.com", company.ExchangeCommissionUrl())
	assert.Equal(t, "http://logo.com", company.LogoUrl())
}
