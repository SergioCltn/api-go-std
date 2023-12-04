package middleware

import (
	"log"
	"net/http"
)

type RateLimiterMiddleware struct {
	next http.Handler
}

func newRateLimiterMiddleware(next http.Handler) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		next: next,
	}
}

func (l *RateLimiterMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("URL called -", r.URL, " ,before calling original handler")
	l.next.ServeHTTP(w, r)
}
