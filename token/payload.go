package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token is expired")
)

// Payload 存储 token 中的信息
type Payload struct {
	ID        uuid.UUID `json:"id"` // 唯一标识 uuid
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issuedAt"`  // 何时创建
	ExpiredAt time.Time `json:"expiredAt"` // 何时失效
}

func NewPayload(username, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
