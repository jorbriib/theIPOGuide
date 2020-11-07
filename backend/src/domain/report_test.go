package domain_test

import (
	. "github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReport(t *testing.T) {
	url := "http://www.url.com"
	message := "this is a message"
	report := NewReport(url, message)

	assert.Equal(t, url, report.Url())
	assert.Equal(t, message, report.Message())
}