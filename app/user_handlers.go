package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tylerchambers/boilerplate/models"
)

type NewUserRequest struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (s *Server) newUserHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req NewUserRequest

		json.NewDecoder(r.Body).Decode(&req)

		u, err := models.NewUser(time.Now(), req.Firstname, req.Lastname, req.Username, req.Email, req.Password)
		if err != nil {
			http.Error(w, "could not register user", http.StatusInternalServerError)
		}
		s.db.Create(u)
		w.Write([]byte(u.ID.String()))
	})
}
