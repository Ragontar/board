package crud

import (
	"account-service/database"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	UserID      string `json:"user_id,omitempty"`
	Email       string `json:"email,omitempty"`
	Credentials string `json:"credentials,omitempty"`
}

type Session struct {
	User   User   `json:"user"`
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

type AuthRequest struct {
	Base64Body      string  `json:"base_64_body"`
	Session         Session `json:"session"`
	Base64Signature string  `json:"base_64_signature"`
}

type AuthResponse struct {
	Authenticated bool
}

type ErrorBadInsertData struct{}
type ErrorAlreadyRegistered struct{}
type ErrorSecretExpiredOrAbsent struct{}
type ErrorRowIsAbsent struct{}

func (e *ErrorBadInsertData) Error() string {
	return "error: corrupted or absent data for insertion"
}

func (e *ErrorRowIsAbsent) Error() string {
	return "error: row is absent"
}

func (e *ErrorAlreadyRegistered) Error() string {
	return "error: such uid/email/credentials already exists"
}

func (e *ErrorSecretExpiredOrAbsent) Error() string {
	return "error: session expired or absent"
}

// GetSecret initializes Secret value from Redis database using Token.
// Returns non-nil error if token is absent.
func (s *Session) GetSecret() error {
	if s.Token == uuid.Nil.String() {
		return fmt.Errorf("error: token is empty")
	}
	var ctx = context.Background()
	var err error

	getRes := database.GetTokenStorage().Get(ctx, s.Token)
	if err = getRes.Err(); err != nil {
		return err
	}

	uuidSecret, err := uuid.Parse(getRes.Val())
	if err != nil {
		return err
	}
	s.Secret = uuidSecret.String()
	setRes := database.GetTokenStorage().Set(ctx, s.Token, s.Secret, redisKeyTTL)

	return setRes.Err()
}

func (r *AuthRequest) CheckPayload() bool {
	if err := r.Session.GetSecret(); err != nil {
		return false
	}

	strToHash := fmt.Sprintf("%s:%s", r.Session.Secret, r.Base64Body)
	hashed := sha256.Sum256([]byte(strToHash))

	return base64.StdEncoding.EncodeToString(hashed[:]) == r.Base64Signature
}
