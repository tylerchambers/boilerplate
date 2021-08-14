package app

import (
	"net/http"
)

// userAuthMiddleware checks if a user is authenticated.
func (s *Server) userAuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, "boilerplate_userauth")
		if err != nil {
			http.Error(w, "Could not decode session info.", http.StatusBadRequest)
			return
		}
		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden - Please login.", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
