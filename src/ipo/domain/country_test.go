package domain_test

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateCountry(t *testing.T) {
	assertion := assert.New(t)
	country := domain.HydrateCountry("1-1", "US", "USA")
	assertion.NotNil(country)

	assertion.Equal(domain.CountryId("1-1"), country.Id())
	assertion.Equal("US", country.Code())
	assertion.Equal("USA", country.Name())
}