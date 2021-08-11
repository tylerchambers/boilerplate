package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Port   string
	Router *mux.Router
	DB     *gorm.DB
	Email  struct{}
}

// initDB initializes the application's database conection.
func (s *Server) initDB() error {
	return nil
}

// initMail initializes the application's mail capibilites.
func (s *Server) initMail() error {
	return nil
}

// NewServer takes a port and returns an instance of the server.
func NewServer(port string) (*Server, error) {
	s := &Server{Port: port}
	err := s.initDB()
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}
	err = s.initMail()
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}
	s.initRoutes()
	return s, nil
}

func (s *Server) Run() {
	addr := fmt.Sprintf(":%s", s.Port)
	log.Printf("Starting app on %s", addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
