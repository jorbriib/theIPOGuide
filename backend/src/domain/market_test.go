package domain_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateMarket(t *testing.T) {
	currency := HydrateCurrency("USD", "American Dollar", "$%s")
	market := HydrateMarket("uuid", "NQ", "Nasdaq", currency, "image", 10)
	assert.NotNil(t, market)

	assert.Equal(t, MarketId("uuid"), market.Id())
	assert.Equal(t, "NQ", market.Code())
	assert.Equal(t, "Nasdaq", market.Name())
	assert.Equal(t, currency, market.Currency())
	assert.Equal(t, "image", market.Image())
	assert.Equal(t, 10, market.TotalIpos())
}