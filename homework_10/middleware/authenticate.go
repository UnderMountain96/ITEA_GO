package middleware

import (
	"fmt"
	"net/http"
)

type AuthenticateMiddleware struct {
	token string
}

func NewAuthenticate(token string) *AuthenticateMiddleware {
	return &AuthenticateMiddleware{token: token}
}

func (m *AuthenticateMiddleware) isValid(userToken string) bool {
	return userToken == fmt.Sprint("Bearer ", m.token)
}

func (m *AuthenticateMiddleware) Wrap(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPatch || r.Method == http.MethodDelete {
			token := r.Header.Get("Authorization")

			if !m.isValid(token) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
