package middleware

import (
	"log"
	"net/http"
)

type LogMiddleware struct {
	next http.Handler
}

func NewLogMiddleware(next http.Handler) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("URL called -", r.URL, " ,before calling original handler")
	l.next.ServeHTTP(w, r)
}
