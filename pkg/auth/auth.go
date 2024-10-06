package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sena1267/cycle-note/config"
	"github.com/sena1267/cycle-note/domain/model"
	"github.com/sena1267/cycle-note/pkg/util/ctxtime"
)

type Authenticator struct {
	AccessTokenExpirationHour int
	AccessTokenSecret         string
}

type customClaims struct {
	jwt.RegisteredClaims
	UserID model.UserID `json:"user_id"`
}

func NewAuthenticator(cfg config.Auth) (Authenticator, error) {
	if cfg.AccessTokenExpirationHour == 0 {
		return Authenticator{}, errors.New("required value: AccessTokenExpirationHour")
	}
	if cfg.AccessTokenSecret == "" {
		return Authenticator{}, errors.New("required value: AccessTokenSecret")
	}

	return Authenticator{
		AccessTokenExpirationHour: cfg.AccessTokenExpirationHour,
		AccessTokenSecret:         cfg.AccessTokenSecret,
	}, nil
}

func (a *Authenticator) CreateAccessToken(ctx context.Context, user model.User) (string, error) {
	claims := customClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ctxtime.Now(ctx).Add(time.Duration(a.AccessTokenExpirationHour) * time.Hour)),
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 署名付きトークンの作成
	signedToken, err := token.SignedString([]byte(a.AccessTokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to create signed token. %w", err)
	}

	return signedToken, err
}

func (a *Authenticator) GetUserIDFromToken(token string) (model.UserID, error) {
	jwtToken, err := jwt.Parse(token, func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %s ", jwtToken.Header["alg"])
		}

		return []byte(a.AccessTokenSecret), nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token. %w", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}

	userID, ok := claims["user_id"]
	if !ok {
		// TODO: エラーメッセージ変更
		return "", errors.New("no user")
	}
	strUserID, ok := userID.(string)
	if !ok {
		// TODO: エラーメッセージ変更
		return "", errors.New("failed to convert")
	}

	return model.UserID(strUserID), nil
}
