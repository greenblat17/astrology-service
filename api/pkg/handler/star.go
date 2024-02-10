package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) getAllStars() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stars, err := h.services.Star.GetAllStars()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stars)
	}
}

func (h *Handler) getStarsByDate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		date := vars["date"]

		star, err := h.services.Star.GetStarByDate(date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(star)
	}
}
