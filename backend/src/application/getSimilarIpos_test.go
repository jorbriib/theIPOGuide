package application_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/application"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGetSimilarIposService(t *testing.T) {
	ir := IpoRepositoryMock{}
	mr := MarketRepositoryMock{}
	cr := CompanyRepositoryMock{}
	service := NewGetSimilarIposService(ir, mr, cr)
	assert.NotNil(t, service)
}

func TestNewGetSimilarIposQuery(t *testing.T) {
	query := NewGetSimilarIposQuery("alias")
	assert.NotNil(t, query)
}
