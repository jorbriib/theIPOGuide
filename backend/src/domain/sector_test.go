package domain_test

import (
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateSector(t *testing.T) {
	assertion := assert.New(t)
	sector := domain.HydrateSector("1-1", "Sector")
	assertion.NotNil(sector)

	assertion.Equal(domain.SectorId("1-1"), sector.Id())
	assertion.Equal("Sector", sector.Name())
}