package domain_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateCountry(t *testing.T) {
	country := HydrateCountry("1-1", "US", "USA")
	assert.NotNil(t, country)

	assert.Equal(t, CountryId("1-1"), country.Id())
	assert.Equal(t, "US", country.Code())
	assert.Equal(t, "USA", country.Name())
}