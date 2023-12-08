package middleware

import (
	"net/http"
)

type AuthenticateMiddleware struct {
	token string
}

func NewAuthenticate(token string) *AuthenticateMiddleware {
	return &AuthenticateMiddleware{token: token}
}

func (m *AuthenticateMiddleware) isValid(userToken string) bool {
	return userToken == m.token
}

func (m *AuthenticateMiddleware) Wrap(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")

		if !m.isValid(authToken) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
