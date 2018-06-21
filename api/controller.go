package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (s *Server) StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "All systems are operating within normal paramaters"}`))
}

func (s *Server) RecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userParam := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recommendations, err := s.Data.generateRecommendations(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(recommendations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
