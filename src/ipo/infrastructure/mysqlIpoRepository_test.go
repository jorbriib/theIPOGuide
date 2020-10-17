package infrastructure_test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/jorbriib/theIPOGuide/src/ipo/infrastructure"
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

	response, err := r.Find()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, domain.IpoId("493506e1-28e2-9e39-8d43-09fdf62ba7dc"), response[0].Id())
}
