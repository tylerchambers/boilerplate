package app

import (
	"encoding/json"
	"net/http"
)

func (s *Server) apiRootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response map[string]string
		json.Unmarshal([]byte(`{ "hello": "world" }`), &response)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusOK)
	}
}
