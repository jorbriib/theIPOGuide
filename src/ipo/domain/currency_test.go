package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateCurrency(t *testing.T) {
	assertion := assert.New(t)
	c := HydrateCurrency("USD", "American Dollar", "$%s")
	assertion.NotNil(c)
}

func TestCurrency_Code(t *testing.T) {
	assertion := assert.New(t)
	c := HydrateCurrency("USD", "American Dollar", "$%s")

	assertion.Equal("USD", c.Code())
}

func TestCurrency_Name(t *testing.T) {
	assertion := assert.New(t)
	c := HydrateCurrency("USD", "American Dollar", "$%s")

	assertion.Equal("American Dollar", c.Name())
}

func TestCurrency_DisplayFromCents(t *testing.T) {
	assertion := assert.New(t)
	c := HydrateCurrency("USD", "American Dollar", "$%s")

	assertion.Equal("$125.53", c.DisplayFromCents(12553))
}