package handler

import (
	"log"
	"net/http"
)

type HomeHandler struct {
	logger *log.Logger
}

func NewHomeHandler(logger *log.Logger) *HomeHandler {
	return &HomeHandler{
		logger: logger,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello!")); err != nil {
		h.logger.Println(err)
	}
}
