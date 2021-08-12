package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) loginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		json.NewDecoder(r.Body).Decode(&req)

		fmt.Fprintf(w, "You tried to login as %s with password %s", req.Email, req.Password)
	}
}
