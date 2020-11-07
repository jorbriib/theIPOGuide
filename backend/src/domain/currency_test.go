package domain_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateCurrency(t *testing.T) {
	c := HydrateCurrency("USD", "American Dollar", "$%s")
	assert.NotNil(t, c)
	assert.Equal(t, "USD", c.Code())
	assert.Equal(t, "American Dollar", c.Name())
	assert.Equal(t, "$125.53", c.DisplayFromCents(12553))
}