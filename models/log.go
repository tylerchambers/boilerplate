package models

import (
	"time"

	"github.com/google/uuid"
)

type LogMessage struct {
	ID      uuid.UUID
	Level   string
	Host    string
	Message string
	Time    time.Time
}
