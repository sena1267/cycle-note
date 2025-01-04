package appctx

import (
	"context"

	"github.com/sena1267/cycle-note/domain/model"
)

type userContextKey struct{}

func ContextWithUser(ctx context.Context, user model.User) context.Context {
	return context.WithValue(ctx, userContextKey{}, user)
}

// User は ctx にセットされたユーザー返す
// セットされていない場合は初期値を返す
func User(ctx context.Context) model.User {
	user, ok := ctx.Value(userContextKey{}).(model.User)
	if !ok {
		return model.User{}
	}
	return user
}
