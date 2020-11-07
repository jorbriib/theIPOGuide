package application_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/application"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAppError(t *testing.T) {
	error := NewAppError("message")
	assert.NotNil(t, error)
	assert.Equal(t, "message", error.Error())
}
