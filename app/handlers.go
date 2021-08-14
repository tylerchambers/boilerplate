package app

import (
	"encoding/json"
	"net/http"
)

func (s *Server) statusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response map[string]string
		json.Unmarshal([]byte(`{ "status": "OK" }`), &response)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) usersOnly() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response map[string]string
		json.Unmarshal([]byte(`{ "msg": "welcome!" }`), &response)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
