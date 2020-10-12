package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHydrateIpo(t *testing.T) {
	assertion := assert.New(t)

	now := time.Now()
	ipo := HydrateIpo("uuid", "uuid", "uuid", 3222, 3444, 10029039, &now)
	assertion.NotNil(ipo)
}

func TestIpo_Id(t *testing.T) {
	assertion := assert.New(t)

	now := time.Now()
	ipo := HydrateIpo("29392-32929da", "uuid", "uuid", 3222, 3444, 10029039, &now)

	assertion.Equal(IpoId("29392-32929da"), ipo.Id())
}

func TestIpo_CompanyId(t *testing.T) {
	assertion := assert.New(t)

	now := time.Now()
	ipo := HydrateIpo("29392-32929da", "m-uuid", "c-uuid", 3222, 3444, 10029039, &now)

	assertion.Equal(CompanyId("c-uuid"), ipo.CompanyId())
}

func TestIpo_MarketId(t *testing.T) {
	assertion := assert.New(t)

	now := time.Now()
	ipo := HydrateIpo("29392-32929da", "m-uuid", "c-uuid", 3222, 3444, 10029039, &now)

	assertion.Equal(MarketId("m-uuid"), ipo.MarketId())
}

func TestIpo_PriceCentsFrom(t *testing.T) {
	assertion := assert.New(t)

	now := time.Now()
	ipo := HydrateIpo("29392-32929da", "m-uuid", "c-uuid", 3222, 3444, 10029039, &now)

	assertion.Equal(uint32(3222), ipo.PriceCentsFrom())
}

func TestIpo_PriceCentsTo(t *testing.T) {
	assertion := assert.New(t)

	now := time.Now()
	ipo := HydrateIpo("29392-32929da", "m-uuid", "c-uuid", 3222, 3444, 10029039, &now)

	assertion.Equal(uint32(3444), ipo.PriceCentsTo())
}

func TestIpo_Shares(t *testing.T) {
	assertion := assert.New(t)

	now := time.Now()
	ipo := HydrateIpo("29392-32929da", "m-uuid", "c-uuid", 3222, 3444, 10029039, &now)

	assertion.Equal(uint32(10029039), ipo.Shares())
}

func TestIpo_ExpectedDate(t *testing.T) {
	assertion := assert.New(t)

	now := time.Now()
	ipo := HydrateIpo("29392-32929da", "m-uuid", "c-uuid", 3222, 3444, 10029039, &now)

	assertion.Equal(&now, ipo.ExpectedDate())
}
