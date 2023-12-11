package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/UnderMountain96/ITEA_GO/http_server/middleware"
	"github.com/UnderMountain96/ITEA_GO/http_server/model"
	"github.com/google/uuid"
)

func TestUserHandler(t *testing.T) {
	staticTime := time.Date(2023, time.December, 11, 12, 0, 0, 0, time.UTC)

	users := []*model.User{
		{
			ID:        uuid.MustParse("b03b5f94-5904-43f2-b1f3-4078c6d47c24"),
			Username:  "user",
			Email:     "user@example.com",
			CreatedAt: staticTime,
			UpdatedAt: staticTime,
		},
		{
			ID:        uuid.MustParse("7d1aeb7b-e248-4c88-be00-8d1332da79e1"),
			Username:  "user 2",
			Email:     "user_2@example.com",
			CreatedAt: staticTime,
			UpdatedAt: staticTime,
		},
	}

	token := "validToken"

	t.Run("Method Get", func(t *testing.T) {
		req := &http.Request{
			Method: http.MethodGet,
		}
		rw := httptest.NewRecorder()

		handler := NewUserHandler(nil)
		handler.users = users
		handler.ServeHTTP(rw, req)

		expectedBody := `[{"id":"b03b5f94-5904-43f2-b1f3-4078c6d47c24","username":"user","email":"user@example.com","update_at":"2023-12-11T12:00:00Z","create_at":"2023-12-11T12:00:00Z"},{"id":"7d1aeb7b-e248-4c88-be00-8d1332da79e1","username":"user 2","email":"user_2@example.com","update_at":"2023-12-11T12:00:00Z","create_at":"2023-12-11T12:00:00Z"}]`

		if strings.Trim(rw.Body.String(), "\n") != expectedBody {
			t.Errorf("invalid response body: got: %s, want: %s", rw.Body, expectedBody)
		}
	})

	t.Run("Method Post", func(t *testing.T) {
		req := &http.Request{
			Method: http.MethodPost,
			Body:   io.NopCloser(bytes.NewBufferString(`{"username": "user 3", "email": "user_3@example.com"}`)),
			Header: http.Header{"Authorization": []string{fmt.Sprint("Bearer ", token)}},
		}
		rw := httptest.NewRecorder()

		authenticateMiddleware := middleware.NewAuthenticate(token)
		handler := NewUserHandler(nil)
		handler.users = users
		authHandler := authenticateMiddleware.Wrap(handler)

		authHandler.ServeHTTP(rw, req)

		expectedCode := http.StatusCreated

		if rw.Code != expectedCode {
			t.Errorf("invalid response status code: got: %d, want: %d", rw.Code, expectedCode)
		}

		expectedUsersLen := 3

		if len(handler.users) != expectedUsersLen {
			t.Errorf("invalid users length: got: %d, want: %d", len(handler.users), expectedUsersLen)
		}
	})

	t.Run("Method Patch", func(t *testing.T) {
		req := &http.Request{
			Method: http.MethodPatch,
			Body:   io.NopCloser(bytes.NewBufferString(`{"id": "b03b5f94-5904-43f2-b1f3-4078c6d47c24", "email": "new_user@example.com"}`)),
			Header: http.Header{"Authorization": []string{fmt.Sprint("Bearer ", token)}},
		}
		rw := httptest.NewRecorder()

		authenticateMiddleware := middleware.NewAuthenticate(token)
		handler := NewUserHandler(nil)
		handler.users = users
		authHandler := authenticateMiddleware.Wrap(handler)

		authHandler.ServeHTTP(rw, req)

		expectedCode := http.StatusOK

		if rw.Code != expectedCode {
			t.Errorf("invalid response status code: got: %d, want: %d", rw.Code, expectedCode)
		}

		user := handler.users[0]

		if user.CreatedAt.Equal(user.UpdatedAt) {
			t.Errorf("UpdatedAt should not be equal CreatedAt")
		}
	})

	t.Run("Method Delete", func(t *testing.T) {
		req := &http.Request{
			Method: http.MethodDelete,
			Body:   io.NopCloser(bytes.NewBufferString(`{"id": "b03b5f94-5904-43f2-b1f3-4078c6d47c24"}`)),
			Header: http.Header{"Authorization": []string{fmt.Sprint("Bearer ", token)}},
		}
		rw := httptest.NewRecorder()

		authenticateMiddleware := middleware.NewAuthenticate(token)
		handler := NewUserHandler(nil)
		handler.users = users
		authHandler := authenticateMiddleware.Wrap(handler)

		authHandler.ServeHTTP(rw, req)

		expectedCode := http.StatusOK

		if rw.Code != expectedCode {
			t.Errorf("invalid response status code: got: %d, want: %d", rw.Code, expectedCode)
		}

		expectedUsersLen := 1

		if len(handler.users) != expectedUsersLen {
			t.Errorf("invalid users length: got: %d, want: %d", len(handler.users), expectedUsersLen)
		}
	})

	t.Run("Invalid Token", func(t *testing.T) {
		req := &http.Request{
			Method: http.MethodPost,
			Body:   io.NopCloser(nil),
			Header: http.Header{"Authorization": []string{fmt.Sprint("Bearer ", "invalidToken")}},
		}
		rw := httptest.NewRecorder()

		authenticateMiddleware := middleware.NewAuthenticate(token)
		authHandler := authenticateMiddleware.Wrap(NewUserHandler(nil))

		authHandler.ServeHTTP(rw, req)

		expectedCode := http.StatusUnauthorized

		if rw.Code != expectedCode {
			t.Errorf("invalid response status code: got: %d, want: %d", rw.Code, expectedCode)
		}
	})
}
