package shared

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserID struct {
	Userid string `json:"user_id"`
}

type UserCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomError struct {
	GivenError  string `json:"error"`
	Description string `json:"error_desc"`
}

// Custom timeout to use instead of context.Background
func OneMinuteContextTimeOut() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	return ctx
}

func GenetareNewUUID(payload UserCreds) (string, error) {
	//Generate new ID from payload (user creds)
	ID, uiderror := uuid.FromBytes([]byte(payload.Username + payload.Password))
	if uiderror != nil {
		return "", uiderror
	}
	return ID.String(), nil
}
