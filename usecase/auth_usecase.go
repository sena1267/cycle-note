package usecase

import (
	"context"
	"fmt"

	"github.com/sena1267/cycle-note/domain/model"
	"github.com/sena1267/cycle-note/domain/repository"
	"github.com/sena1267/cycle-note/pkg/auth"
	"github.com/sena1267/cycle-note/pkg/xidgen"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repo          repository.UserRepository
	authenticator auth.Authenticator
}

func NewAuthUsecase(repo repository.UserRepository) Auth {
	return Auth{repo: repo}
}

type SignUpInput struct {
	Email    string
	Name     string
	Password string
}

type SignUpOutput struct {
	Token string
}

type SignInInput struct {
	Email    string
	Password string
}

type SignInOutput struct {
	Token string
}

func (uc *Auth) SignUp(ctx context.Context, input SignUpInput) (SignUpOutput, error) {
	hashPassword, err := encryptPassword(input.Password)
	if err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to encrypt password. %s", err)
	}

	newUser := model.User{
		ID:       model.UserID(xidgen.GenerateXID()),
		Name:     input.Name,
		Password: hashPassword,
		Email:    input.Email,
	}
	if err = uc.repo.Create(ctx, newUser); err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to create user. %w", err)
	}

	// TODO: authenticator の初期化を別の場所で行う
	authenticator := auth.Authenticator{
		AccessTokenExpirationHour: 12,
		AccessTokenSecret:         "secret",
	}

	token, err := authenticator.CreateAccessToken(ctx, newUser)
	if err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to create access token. %w", err)
	}

	return SignUpOutput{Token: token}, nil
}

func (uc *Auth) SignIn(ctx context.Context, input SignInInput) (SignInOutput, error) {
	user, err := uc.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return SignInOutput{}, fmt.Errorf("failed to get user by email. %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return SignInOutput{}, fmt.Errorf("invalid password. %w", err)
	}

	token, err := uc.authenticator.CreateAccessToken(ctx, *user)
	if err != nil {
		return SignInOutput{}, fmt.Errorf("failed to create access token. %w", err)
	}

	userID, _ := uc.authenticator.GetUserIDFromToken(token)
	fmt.Println(userID)

	return SignInOutput{Token: token}, nil

}

// TODO: 関数の定義場所を考える
func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password. %w", err)
	}
	return string(hash), nil
}
