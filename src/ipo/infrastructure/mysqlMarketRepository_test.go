package infrastructure_test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/jorbriib/theIPOGuide/src/ipo/infrastructure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMySQLMarketRepository(t *testing.T) {
	r := infrastructure.NewMySQLMarketRepository(db)
	assert.NotNil(t, r)
}

func TestMySQLMarketRepository_GetById(t *testing.T) {
	r := infrastructure.NewMySQLMarketRepository(db)

	id := domain.MarketId("47a0bb30-a9da-11f6-9f27-b52510f1cc6a")

	response, err := r.GetById(id)

	assert.Nil(t, err)
	assert.Equal(t, id, response.Id())
}

func TestMySQLMarketRepository_GetById_ReturnsNilWhenNotFound(t *testing.T) {
	r := infrastructure.NewMySQLMarketRepository(db)

	id := domain.MarketId("1293f9f9-c2b7-1e7b-8271-77a4ce70c6f0")

	response, err := r.GetById(id)

	assert.Nil(t, err)
	assert.Nil(t, response)
}

func TestMySQLMarketRepository_FindByIds_ReturnsSliceLength0_WhenNoMarketIds(t *testing.T) {
	r := infrastructure.NewMySQLMarketRepository(db)

	var ids []domain.MarketId
	response, err := r.FindByIds(ids)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(response))
}

func TestMySQLMarketRepository_FindByIds(t *testing.T) {
	r := infrastructure.NewMySQLMarketRepository(db)

	ids := []domain.MarketId{
		domain.MarketId("47a0bb30-a9da-11f6-9f27-b52510f1cc6a"),
	}
	response, err := r.FindByIds(ids)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, ids[0], response[0].Id())
}
