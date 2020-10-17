package domain_test

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateCountry(t *testing.T) {
	assertion := assert.New(t)
	country := domain.HydrateCountry("US", "USA")
	assertion.NotNil(country)
}

func TestCountry_Symbol(t *testing.T) {
	assertion := assert.New(t)
	country := domain.HydrateCountry("US", "USA")

	assertion.Equal("US", country.Code())
}

func TestCountry_Name(t *testing.T) {
	assertion := assert.New(t)
	country := domain.HydrateCountry("US", "USA")

	assertion.Equal("USA", country.Name())
}