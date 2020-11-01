package api_test

import (
	"github.com/jorbriib/theIPOGuide/backend/src/ui/public/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContentTypeMiddleware(t *testing.T) {
	handler := api.ContentTypeMiddleware(
		"my-type",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ty := w.Header().Get("Content-Type")
			assert.Equal(t, "my-type", ty)
		}),
	)

	r := &http.Request{}
	w := &httptest.ResponseRecorder{}
	handler.ServeHTTP(w, r)
}

func TestEnableCorsMiddleware(t *testing.T) {
	handler := api.EnableCorsMiddleware(
		"origin",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			or := w.Header().Get("Access-Control-Allow-Origin")
			assert.Equal(t, "origin", or)
		}),
	)

	r := &http.Request{}
	w := &httptest.ResponseRecorder{}
	handler.ServeHTTP(w, r)
}
