package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateCompany(t *testing.T) {
	assertion := assert.New(t)
	company := HydrateCompany("PINS", "Pinterest")
	assertion.NotNil(company)
}

func TestCompany_Symbol(t *testing.T) {
	assertion := assert.New(t)
	company := HydrateCompany("PINS", "Pinterest")

	assertion.Equal("PINS", company.Symbol())
}

func TestCompany_Name(t *testing.T) {
	assertion := assert.New(t)
	company := HydrateCompany("PINS", "Pinterest")

	assertion.Equal("Pinterest", company.Name())
}