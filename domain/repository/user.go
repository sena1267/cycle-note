package repository

import (
	"context"

	"github.com/sena1267/cycle-note/domain/model"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, user model.User) error
}
