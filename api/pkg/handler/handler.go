package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/greenblat17/api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/stars", h.getAllStars()).Methods(http.MethodGet)
	r.HandleFunc("/stars/{date}", h.getStarsByDate()).Methods(http.MethodGet)

	return r
}
