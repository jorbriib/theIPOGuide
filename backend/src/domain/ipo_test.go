package domain_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHydrateIpo(t *testing.T) {
	

	now := time.Now()
	ipo := HydrateIpo("uuid",  "alias", "intro", "m-uuid", "c-uuid", 3222, 3444, 10029039, &now)
	assert.NotNil(t, ipo)

	assert.Equal(t, IpoId("uuid"), ipo.Id())
	assert.Equal(t, "alias", ipo.Alias())
	assert.Equal(t, "intro", ipo.Intro())
	assert.Equal(t, MarketId("m-uuid"), ipo.MarketId())
	assert.Equal(t, CompanyId("c-uuid"), ipo.CompanyId())

	assert.Equal(t, uint32(3222), ipo.PriceCentsFrom())
	assert.Equal(t, uint32(3444), ipo.PriceCentsTo())
	assert.Equal(t, uint32(10029039), ipo.Shares())
	assert.Equal(t, &now, ipo.ExpectedDate())
}
