package utils

import (
	"time"
)

type SessionData struct {
	ID       int
	Username string
        Role string
	Expiry   time.Time
}

type ContextKey string
