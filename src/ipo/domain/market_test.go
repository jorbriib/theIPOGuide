package domain_test

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateMarket(t *testing.T) {
	assertion := assert.New(t)
	currency := domain.HydrateCurrency("USD", "American Dollar", "$%s")
	market := domain.HydrateMarket("uuid", "NQ", "Nasdaq", currency)
	assertion.NotNil(market)
}

func TestMarket_Id(t *testing.T) {
	assertion := assert.New(t)
	currency := domain.HydrateCurrency("USD", "American Dollar", "$%s")
	market := domain.HydrateMarket("uuid", "NQ", "Nasdaq", currency)

	assertion.Equal(domain.MarketId("uuid"), market.Id())
}

func TestMarket_Code(t *testing.T) {
	assertion := assert.New(t)
	currency := domain.HydrateCurrency("USD", "American Dollar", "$%s")
	market := domain.HydrateMarket("uuid", "NQ", "Nasdaq", currency)

	assertion.Equal("NQ", market.Code())
}

func TestMarket_Name(t *testing.T) {
	assertion := assert.New(t)
	currency := domain.HydrateCurrency("USD", "American Dollar", "$%s")
	market := domain.HydrateMarket("uuid", "NQ", "Nasdaq", currency)

	assertion.Equal("Nasdaq", market.Name())
}


func TestMarket_Currency(t *testing.T) {
	assertion := assert.New(t)
	currency := domain.HydrateCurrency("USD", "American Dollar", "$%s")
	market := domain.HydrateMarket("uuid", "NQ", "Nasdaq", currency)

	assertion.Equal(currency, market.Currency())
}