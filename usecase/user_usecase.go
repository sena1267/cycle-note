package usecase

import (
	"context"

	"github.com/sena1267/cycle-note/domain/model"
	"github.com/sena1267/cycle-note/domain/repository"
)

type User struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) User {
	return User{userRepo: repo}
}

type MeInput struct {
	token string
}

type MeOutput struct {
	user model.User
}

func (uc *User) Me(ctx context.Context, input MeInput) (MeOutput, error) {
	return MeOutput{}, nil
}
