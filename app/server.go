package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/tylerchambers/boilerplate/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	port         string
	router       *mux.Router
	db           *gorm.DB
	email        struct{}
	secretKey    []byte
	sessionStore *sessions.CookieStore
}

// initDB initializes the application's database conection.
func (s *Server) initDB() error {
	var err error
	s.db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to connect to the database")
	}
	return nil
}

// migrateDB performs DB migrations.
func (s *Server) migrateDB() error {
	s.db.AutoMigrate(&models.User{})
	// TODO: Remove test user creation.
	testUser, _ := models.NewUser(time.Now(), "test", "test@example.com", "password1")
	s.db.Create(testUser)

	return nil
}

// initMail initializes the application's mail capibilites.
func (s *Server) initMail() error {
	s.email = struct{}{}
	return nil
}

// initSecretKey takes a key from an SECRET_KEY env variable, checks its size, and sets for the server.
func (s *Server) initSecretKey() error {
	key := os.Getenv("SECRET_KEY")
	bkey := []byte(key)
	keySize := len(bkey)
	if keySize == 32 || keySize == 24 || keySize == 16 {
		s.secretKey = bkey
		return nil
	}
	return errors.New("secret key could not be set - only 32, 24 and 16 byte long keys accepted")
}

// initSessionStore sets up a session store using the server's secret key.
func (s *Server) initSessionStore() error {
	s.sessionStore = sessions.NewCookieStore(s.secretKey)
	return nil
}

// NewServer takes a port and returns an instance of the server.
func NewServer(port string) (*Server, error) {
	s := &Server{port: port}
	err := s.initDB()
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}
	s.migrateDB()
	err = s.initSecretKey()
	if err != nil {
		return nil, err
	}
	_ = s.initSessionStore()
	err = s.initMail()
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}
	s.initRoutes()
	return s, nil
}

func (s *Server) Run() {
	addr := fmt.Sprintf(":%s", s.port)
	log.Printf("Starting app on %s", addr)
	log.Fatal(http.ListenAndServe(addr, s.router))
}
