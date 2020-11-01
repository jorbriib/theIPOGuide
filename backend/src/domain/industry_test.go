package domain_test

import (
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateIndustry(t *testing.T) {
	assertion := assert.New(t)
	sector := domain.HydrateIndustry("1-1", "Industry")
	assertion.NotNil(sector)

	assertion.Equal(domain.IndustryId("1-1"), sector.Id())
	assertion.Equal("Industry", sector.Name())
}