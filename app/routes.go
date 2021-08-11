package app

import "github.com/gorilla/mux"

func (s *Server) initRoutes() {
	s.Router = mux.NewRouter()

	apiSubrouter := s.Router.PathPrefix("/api").Subrouter()
	// TODO: Enable middleware here.
	// apiSubrouter.Use()
	apiSubrouter.HandleFunc("/", s.apiRootHandler())
}
