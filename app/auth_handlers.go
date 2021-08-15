package app

import (
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"
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

		session, _ := s.sessionStore.Get(r, "boilerplate_userauth")

		var user models.User
		s.db.First(&user, "email = ?", req.Email)
		spew.Dump(user)
		if user.CheckPassword(req.Password) {
			session.Values["authenticated"] = true
			session.Values["user_id"] = user.ID.String()
			session.Save(r, w)
			w.Write([]byte("Login successful!"))
		} else {
			w.Write([]byte("Login failed."))
		}
	}
}

func (s *Server) logoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.sessionStore.Get(r, "boilerplate_userauth")

		session.Values["authenticated"] = false
		session.Save(r, w)
		w.Write([]byte("Logout successful"))
	}
}
