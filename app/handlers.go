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

type UsersOnlyResponse struct {
	Uid string `json:"uid,omitempty"`
}

func (s *Server) usersOnly() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.sessionStore.Get(r, "boilerplate_userauth")
		uid, OK := session.Values["user_id"].(string)
		if !OK {
			http.Error(w, "Could not decode session info.", http.StatusBadRequest)
		}
		resp := &UsersOnlyResponse{Uid: uid}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
