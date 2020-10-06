package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateCountry(t *testing.T) {
	assertion := assert.New(t)
	country := HydrateCountry("US", "USA")
	assertion.NotNil(country)
}

func TestCountry_Symbol(t *testing.T) {
	assertion := assert.New(t)
	country := HydrateCountry("US", "USA")

	assertion.Equal("US", country.Code())
}

func TestCountry_Name(t *testing.T) {
	assertion := assert.New(t)
	country := HydrateCountry("US", "USA")

	assertion.Equal("USA", country.Name())
}