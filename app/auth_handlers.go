package app

import (
	"encoding/json"
	"net/http"

	"github.com/tylerchambers/boilerplate/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) loginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		json.NewDecoder(r.Body).Decode(&req)

		var user models.User
		s.db.First(&user, "email = ?", req.Email)
		if user.CheckPassword(req.Password) {
			w.Write([]byte("login successful!"))
		} else {
			w.Write([]byte("login failed"))
		}
	}
}
