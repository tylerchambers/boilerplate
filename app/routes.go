package app

import "github.com/gorilla/mux"

func (s *Server) initRoutes() {
	s.router = mux.NewRouter()

	// A subrouter to handle API requests.
	apiSubrouter := s.router.PathPrefix("/api/v1").Subrouter()
	apiSubrouter.HandleFunc("/status", s.statusHandler())

	// Authentication related routes.
	apiSubrouter.HandleFunc("/auth/login", s.loginHandler()).Methods("POST")

	apiSubrouter.HandleFunc("/usersonly", s.userAuthMiddleware(s.usersOnly()))
}
