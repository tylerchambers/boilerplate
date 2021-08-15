package app

import (
	"encoding/json"
	"fmt"
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

type UsersOnlyResponse struct {
	Uid string `json:"uid,omitempty"`
}

func (s *Server) usersOnly() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := fmt.Sprintf("%v", (r.Context().Value(uidString("uid"))))
		resp := &UsersOnlyResponse{Uid: fmt.Sprintf("%v", uid)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
