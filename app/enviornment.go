package app

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Eviornment stores our enviornment variables.
type Enviornment struct {
	Port          string
	Debug         bool
	SecretKey     string
	PgHost        string
	PgPort        string
	PgUser        string
	PgPassword    string
	PgDBName      string
	PgSSLMode     string
	PgTZ          string
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
	MailDomain    string
	MailAPIKey    string
}

func (s *Server) initEnv() {
	var e Enviornment
	err := envconfig.Process("bpapp", &e)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.env = &e
}
