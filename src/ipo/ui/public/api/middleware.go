package api

import (
	"golang.org/x/time/rate"
	"net/http"
)

func ContentTypeMiddleware(contentType string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)

		next.ServeHTTP(w, r)
	})
}

func EnableCorsMiddleware(origin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)

		next.ServeHTTP(w, r)
	})
}

var limiter *rate.Limiter
func ThrottleMiddleware(limit float64, bucket int, next http.Handler) http.Handler {
	limiter = rate.NewLimiter(rate.Limit(limit), bucket)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}