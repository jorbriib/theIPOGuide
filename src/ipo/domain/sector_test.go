package domain_test

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateSector(t *testing.T) {
	assertion := assert.New(t)
	sector := domain.HydrateSector("Sector")
	assertion.NotNil(sector)
}

func TestSector_Name(t *testing.T) {
	assertion := assert.New(t)
	sector := domain.HydrateSector("Sector")

	assertion.Equal("Sector", sector.Name())
}