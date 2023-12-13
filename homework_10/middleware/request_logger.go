package middleware

import (
	"log"
	"net/http"
	"time"
)

type RequestLoggerMiddleware struct {
	logger *log.Logger
}

func NewRequestLogger(logger *log.Logger) *RequestLoggerMiddleware {
	return &RequestLoggerMiddleware{logger: logger}
}

func (m *RequestLoggerMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))

	})
}
