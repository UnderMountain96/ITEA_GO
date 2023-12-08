package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/UnderMountain96/ITEA_GO/http_server/model"
	"github.com/google/uuid"
)

type UserHandler struct {
	logger *log.Logger
	users  []*model.User
}

func NewUserHandler(logger *log.Logger) *UserHandler {
	return &UserHandler{
		logger: logger,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	switch r.Method {
	case http.MethodGet:
		if err := h.handleGet(w, r); err != nil {
			http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		}
	case http.MethodPost:
		if err := h.handlePost(w, r); err != nil {
			http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		}
	case http.MethodPatch:
		if err := h.handlePatch(w, r); err != nil {
			http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		}
	case http.MethodDelete:
		if err := h.handleDelete(w, r); err != nil {
			http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		}
	default:
		http.Error(w, `{"error": "invalid HTTP method"}`, http.StatusBadRequest)
	}
}

func (h *UserHandler) handleGet(w http.ResponseWriter, r *http.Request) error {
	return json.NewEncoder(w).Encode(h.users)
}

func (h *UserHandler) handlePost(w http.ResponseWriter, r *http.Request) error {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	req := &request{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, req); err != nil {
		return err
	}

	user := model.NewUser(req.Username, req.Email)

	h.users = append(h.users, user)

	w.WriteHeader(http.StatusCreated)

	return nil
}

func (h *UserHandler) handlePatch(w http.ResponseWriter, r *http.Request) error {
	type request struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
	}
	req := &request{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, req); err != nil {
		return err
	}

	for _, u := range h.users {
		if u.ID == req.ID {
			if req.Username != "" {
				u.Username = req.Username
			}
			if req.Email != "" {
				u.Email = req.Email
			}
			u.UpdateAt = time.Now()

			w.WriteHeader(http.StatusOK)

			return nil
		}
	}

	w.WriteHeader(http.StatusBadRequest)

	return nil
}

func (h *UserHandler) handleDelete(w http.ResponseWriter, r *http.Request) error {
	type request struct {
		ID uuid.UUID `json:"id"`
	}
	req := &request{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, req); err != nil {
		return err
	}

	for i, u := range h.users {
		if u.ID == req.ID {
			h.users = append(h.users[:i], h.users[i+1:]...)

			w.WriteHeader(http.StatusCreated)

			return nil
		}
	}

	w.WriteHeader(http.StatusBadRequest)

	return nil
}
