package domain_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateIndustry(t *testing.T) {
	sector := HydrateIndustry("1-1", "industry", "Industry")
	assert.NotNil(t, sector)

	assert.Equal(t, IndustryId("1-1"), sector.Id())
	assert.Equal(t, "industry", sector.Alias())
	assert.Equal(t, "Industry", sector.Name())
}