package api

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ServiceMock struct{
	mock.Mock
}

func (s ServiceMock) GetIPOs(query application.GetIposQuery) (*application.GetIposResponse, error) {
	args := s.Called(query)
	return args.Get(0).(*application.GetIposResponse), args.Error(1)
}

func TestNewController(t *testing.T) {
	assertion := assert.New(t)

	s := ServiceMock{}
	service := NewController(s)

	assertion.NotNil(service)
}

func TestController_GetIpos(t *testing.T) {
	assertion := assert.New(t)

	s := ServiceMock{}
	query := application.NewGetIposQuery()
	expectedResponse := &application.GetIposResponse{}
	s.On("GetIPOs", query).Return(expectedResponse, nil)

	service := NewController(s)
	w := httptest.NewRecorder()
	service.GetIpos(w, &http.Request{})

	assertion.Equal("[]\n", w.Body.String())
}