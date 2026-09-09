package main

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// Global rate limiter
var limiter = rate.NewLimiter(rate.Every(time.Second), 10) // 10 requests per second

func globalLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", globalLimitMiddleware(helloHandler))
	http.ListenAndServe(":8080", nil)
}
