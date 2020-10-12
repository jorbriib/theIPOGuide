package infrastructure

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMySQLMarketRepository(t *testing.T) {
	r := NewMySQLMarketRepository(db)
	assert.NotNil(t, r)
}

func TestMySQLMarketRepository_FindByIds_ReturnsSliceLength0_WhenNoMarketIds(t *testing.T) {
	r := NewMySQLMarketRepository(db)

	var ids []domain.MarketId
	response, err := r.FindByIds(ids)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(response))
}

func TestMySQLMarketRepository_FindByIds(t *testing.T) {
	r := NewMySQLMarketRepository(db)

	ids := []domain.MarketId{
		domain.MarketId("47a0bb30-a9da-11f6-9f27-b52510f1cc6a"),
	}
	response, err := r.FindByIds(ids)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, ids[0], response[0].Id())
}
