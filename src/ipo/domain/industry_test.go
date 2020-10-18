package domain_test

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateIndustry(t *testing.T) {
	assertion := assert.New(t)
	sector := domain.HydrateIndustry("Industry")
	assertion.NotNil(sector)
}

func TestIndustry_Name(t *testing.T) {
	assertion := assert.New(t)
	sector := domain.HydrateIndustry("Industry")

	assertion.Equal("Industry", sector.Name())
}