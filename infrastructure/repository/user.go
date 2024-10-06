package repository

import (
	"context"
	"fmt"

	"github.com/sena1267/cycle-note/domain/model"
	"github.com/sena1267/cycle-note/domain/repository"
	"github.com/sena1267/cycle-note/infrastructure"
)

type UserRepository struct {
	dbClient *infrastructure.DBClient
}

func NewUserRepository(dbClient *infrastructure.DBClient) repository.UserRepository {
	return &UserRepository{dbClient: dbClient}
}

func (u *UserRepository) Create(ctx context.Context, user model.User) error {
	_, err := u.dbClient.DB.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to insert user. %w", err)
	}

	return nil
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	if err := u.dbClient.DB.NewSelect().Model(user).Where("email = ?", email).Scan(ctx); err != nil {
		return nil, fmt.Errorf("user not found. %w", err)
	}

	return user, nil
}
