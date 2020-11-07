package infrastructure_test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/jorbriib/theIPOGuide/backend/src/infrastructure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMySQLIpoRepository(t *testing.T) {
	r := infrastructure.NewMySQLIpoRepository(db)
	assert.NotNil(t, r)
}

func TestMySQLIpoRepository_GetById(t *testing.T) {
	r := infrastructure.NewMySQLIpoRepository(db)

	alias := "pinterest"

	response, err := r.GetByAlias(alias)

	assert.Nil(t, err)
	assert.Equal(t, alias, response.Alias())
}

func TestMySQLIpoRepository_GetById_ReturnsNilNotFound(t *testing.T) {
	r := infrastructure.NewMySQLIpoRepository(db)

	alias := "alias-not-found"

	response, err := r.GetByAlias(alias)

	assert.Nil(t, err)
	assert.Nil(t, response)
}

func TestMySQLIpoRepository_Find(t *testing.T) {
	r := infrastructure.NewMySQLIpoRepository(db)

	response, err := r.Find( []domain.MarketId{}, []domain.CountryId{}, []domain.SectorId{}, []domain.IndustryId{}, []domain.IpoId{}, "", 0, 20)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(response))
	assert.Equal(t, domain.IpoId("4de1f713-410a-d8b5-b67d-3658e4a89723"), response[0].Id())
}
