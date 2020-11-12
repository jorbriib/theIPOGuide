package domain_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateSector(t *testing.T) {
	sector := HydrateSector("1-1",  "sector", "Sector", "image", 10)
	assert.NotNil(t, sector)

	assert.Equal(t, SectorId("1-1"), sector.Id())
	assert.Equal(t, "sector", sector.Alias())
	assert.Equal(t, "Sector", sector.Name())
	assert.Equal(t, "image", sector.Image())
	assert.Equal(t, 10, sector.TotalIpos())
}