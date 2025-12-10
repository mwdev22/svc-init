package handlers

import "net/http"

type Handler struct {
	// add your dependencies here, e.g., services, caches, etc.
}

func NewHandler() *Handler {
	return &Handler{
		// initialize your dependencies here
	}
}

func (h *Handler) GetExample(w http.ResponseWriter, r *http.Request) {
	// handle GET request
	w.Write([]byte("GET example"))
}
