package api

import (
	"encoding/json"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

// GoogleRecaptchaResponse is the google recaptcha response
type GoogleRecaptchaResponse struct {
	Success            bool     `json:"success"`
	ChallengeTimestamp string   `json:"challenge_ts"`
	Hostname           string   `json:"hostname"`
	ErrorCodes         []string `json:"error-codes"`
}

// VerifyRecaptcha verifies the Captcha from the coming requests
func VerifyRecaptcha(recaptchaUrl string, secret string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.FormValue("token")
		if token == "" {
			http.Error(w, http.StatusText(403), http.StatusForbidden)
			return
		}

		req, err := http.PostForm(recaptchaUrl, url.Values{
			"secret":   {secret},
			"response": {token},
		})
		if err != nil {
			log.Println("error issuing the POST to: " + recaptchaUrl)
			log.Println(err)
			http.Error(w, http.StatusText(403), http.StatusForbidden)
			return
		}
		defer func() { req.Body.Close() }()

		body, err := ioutil.ReadAll(req.Body) // Read the response from Google
		if err != nil {
			http.Error(w, http.StatusText(403), http.StatusForbidden)
			return
		}

		var googleResponse GoogleRecaptchaResponse
		err = json.Unmarshal(body, &googleResponse) // Parse the JSON response from Google
		if err != nil {
			log.Println("error unmarshalling the response")
			log.Println(err)
			http.Error(w, http.StatusText(403), http.StatusForbidden)
			return
		}
		if !googleResponse.Success {
			log.Println("no success recaptcha process")
			log.Println(googleResponse)
			http.Error(w, http.StatusText(403), http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
