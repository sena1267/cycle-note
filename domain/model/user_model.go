package model

import (
	"context"
	"errors"

	"github.com/uptrace/bun"
)

type UserID string

type User struct {
	bun.BaseModel `bun:"table:user"`
	ID            UserID
	Name          string
	Email         string
	Password      string
	TimeStamp
}

type userContextKey struct{}

// ContextWithUser -
// TODO: 定義場所検討
func ContextWithUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userContextKey{}, user)
}

func GetUserFromContext(ctx context.Context) (User, error) {
	user, ok := ctx.Value(userContextKey{}).(User)
	if !ok {
		return user, errors.New("no user is associated with key")
	}
	return user, nil
}
