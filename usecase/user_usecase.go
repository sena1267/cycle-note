package usecase

import (
	"github.com/sena1267/cycle-note/domain/repository"
)

type UserUsecase interface{}

type userUsecase struct {
	userRepo repository.UserRepository
}

type (
	UserCreateInput struct {
		Name     string
		Email    string
		Password string
	}
)

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return userUsecase{userRepo: repo}
}
