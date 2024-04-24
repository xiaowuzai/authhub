package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	CreateToken(usernames, role string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
