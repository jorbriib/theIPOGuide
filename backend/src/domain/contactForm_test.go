package domain_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContactForm(t *testing.T) {

	contactForm := NewContactForm("name", "email", "message")

	assert.Equal(t, "name", contactForm.Name())
	assert.Equal(t, "email", contactForm.Email())
	assert.Equal(t, "message", contactForm.Message())
}