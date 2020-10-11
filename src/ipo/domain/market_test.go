package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateMarket(t *testing.T) {
	assertion := assert.New(t)
	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)
	assertion.NotNil(market)
}

func TestMarket_Symbol(t *testing.T) {
	assertion := assert.New(t)
	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)

	assertion.Equal("NQ", market.Code())
}

func TestMarket_Name(t *testing.T) {
	assertion := assert.New(t)
	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)

	assertion.Equal("Nasdaq", market.Name())
}


func TestMarket_Currency(t *testing.T) {
	assertion := assert.New(t)
	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("NQ", "Nasdaq", currency)

	assertion.Equal(currency, market.Currency())
}