package application

import (
	"errors"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type IpoRepositoryMock struct{
	mock.Mock
}

func (r IpoRepositoryMock) Find() ([]domain.Ipo, error){
	args := r.Called()
	return args.Get(0).([]domain.Ipo), args.Error(1)
}

func TestNewService(t *testing.T) {
	assertion := assert.New(t)
	r := IpoRepositoryMock{}
	service := NewService(r)
	assertion.NotNil(service)
}

func TestService_GetIPOs_FailsWhenRepositoryReturnsError(t *testing.T) {
	assertion := assert.New(t)

	r := IpoRepositoryMock{}
	r.On("Find").Return([]domain.Ipo{}, errors.New("repository error"))

	service := NewService(r)
	query := NewGetIposQuery()
	ipos, err := service.GetIPOs(query)

	assertion.Equal(0, len(ipos))
	assertion.NotNil(err)
}

func TestService_GetIPOs(t *testing.T) {
	assertion := assert.New(t)

	expectedReturn := []domain.Ipo{
		domain.Ipo{},
		domain.Ipo{},
	}

	r := IpoRepositoryMock{}
	r.On("Find").Return(expectedReturn, nil)

	service := NewService(r)
	query := NewGetIposQuery()
	ipos, err := service.GetIPOs(query)

	assertion.Equal(2, len(ipos))
	assertion.Nil(err)
}

func TestNewGetIposQuery(t *testing.T) {
	assertion := assert.New(t)
	query := NewGetIposQuery()
	assertion.NotNil(query)
}