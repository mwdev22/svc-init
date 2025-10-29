package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mwdev22/rest/middleware"
)

func (a *Api) Mount(r chi.Router) {
	r.Use(middleware.Logger, middleware.RealIP)
	r.Get("/ping", a.Ping)
}

// ping example
// @Summary     Ping the server
// @Description Simple endpoint to check if the server is running
// @Accept      json
// @Produce     json
// @Router      /ping [get]
func (a *Api) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"pong"}`))
}
