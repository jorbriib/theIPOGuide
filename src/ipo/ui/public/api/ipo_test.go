package api_test

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	ipo_public_api "github.com/jorbriib/theIPOGuide/src/ipo/ui/public/api"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestNewController(t *testing.T) {
	assertion := assert.New(t)

	s := application.IpoService{}
	service := ipo_public_api.NewController(s)

	assertion.IsType(ipo_public_api.Controller{}, service)
}
