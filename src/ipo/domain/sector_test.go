package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHydrateSector(t *testing.T) {
	assertion := assert.New(t)
	sector := HydrateSector("Industry")
	assertion.NotNil(sector)
}

func TestSector_Name(t *testing.T) {
	assertion := assert.New(t)
	sector := HydrateSector("Industry")

	assertion.Equal("Industry", sector.Name())
}