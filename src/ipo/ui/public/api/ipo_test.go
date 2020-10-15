package api

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type ServiceMock struct{
	mock.Mock
}

func (s ServiceMock) GetIPOs(query application.GetIposQuery) (*application.GetIposResponse, error) {
	args := s.Called(query)
	return args.Get(0).(*application.GetIposResponse), args.Error(1)
}

func (s ServiceMock) GetIPO(query application.GetIpoQuery) (*application.GetIpoResponse, error) {
	args := s.Called(query)
	return args.Get(0).(*application.GetIpoResponse), args.Error(1)
}

func TestNewController(t *testing.T) {
	assertion := assert.New(t)

	s := ServiceMock{}
	service := NewController(s)

	assertion.NotNil(service)
}
