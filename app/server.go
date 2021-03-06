package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tylerchambers/boilerplate/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	router    *mux.Router
	db        *gorm.DB
	email     struct{}
	secretKey []byte
	env       *Enviornment
}

// initDB initializes the application's database conection.
func (s *Server) initDB() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		s.env.PgHost, s.env.PgUser, s.env.PgPassword, s.env.PgDBName, s.env.PgPort, s.env.PgSSLMode, s.env.PgTZ)
	s.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to connect to the database")
	}
	return nil
}

// migrateDB performs DB migrations.
func (s *Server) migrateDB() error {
	s.db.AutoMigrate(&models.User{})

	return nil
}

// initMail initializes the application's mail capibilites.
func (s *Server) initMail() error {
	s.email = struct{}{}
	return nil
}

// initSecretKey takes a key from an SECRET_KEY env variable, checks its size, and sets for the server.
func (s *Server) initSecretKey() error {
	key := s.env.SecretKey
	bkey := []byte(key)
	keySize := len(bkey)
	if keySize == 32 || keySize == 24 || keySize == 16 {
		s.secretKey = bkey
		return nil
	}
	return errors.New("secret key could not be set - only 32, 24 and 16 byte long keys accepted")
}

// NewServer returns an instance of the server.
func NewServer() (*Server, error) {
	s := &Server{}
	s.initEnv()
	err := s.initDB()
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}
	s.migrateDB()
	err = s.initSecretKey()
	if err != nil {
		return nil, err
	}
	err = s.initMail()
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}
	s.initRoutes()
	return s, nil
}

func (s *Server) Run() {
	addr := fmt.Sprintf(":%s", s.env.Port)
	log.Printf("Starting app on %s", addr)
	log.Fatal(http.ListenAndServe(addr, s.router))
}
