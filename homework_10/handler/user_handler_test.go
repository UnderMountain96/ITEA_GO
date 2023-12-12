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
		expectedUser := model.User{
			Username: "user 3",
			Email:    "user_3@example.com",
		}
		req := &http.Request{
			Method: http.MethodPost,
			Body: io.NopCloser(
				bytes.NewBufferString(
					fmt.Sprintf(`{"username": %q, "email": %q}`, expectedUser.Username, expectedUser.Email),
				),
			),
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

		user := &model.User{}

		for _, u := range handler.users {
			if u.Username == expectedUser.Username {
				user = u
			}
		}

		if user.Username != expectedUser.Username {
			t.Errorf("invalid Username: got: %s, want: %s", user.Username, expectedUser.Username)
		}

		if user.Email != expectedUser.Email {
			t.Errorf("invalid Email: got: %s, want: %s", user.Email, expectedUser.Email)
		}

		if !user.CreatedAt.Equal(user.UpdatedAt) {
			t.Errorf("UpdatedAt must be equal CreatedAt: got: %s, want: %s", user.UpdatedAt, user.CreatedAt)
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
		expectedId := uuid.MustParse("b03b5f94-5904-43f2-b1f3-4078c6d47c24")
		req := &http.Request{
			Method: http.MethodDelete,
			Body: io.NopCloser(
				bytes.NewBufferString(
					fmt.Sprintf(`{"id": %q}`, expectedId),
				),
			),
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

		for _, u := range handler.users {
			if u.ID == expectedId {
				t.Errorf("invalid user with ID: %q must be deleted", expectedId)
			}
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
