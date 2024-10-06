package handler

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	authv1 "github.com/sena1267/cycle-note/gen/protobuf/auth/v1"
	"github.com/sena1267/cycle-note/gen/protobuf/auth/v1/authv1connect"
	"github.com/sena1267/cycle-note/usecase"
)

// TODO: おまじないの扱い
// こんな「おまじない」があるらしい
// 未実装のメソッドを見つけるのに便利だが、NewAuthHandler でもエラー出るんだよなあ
var _ authv1connect.AuthServiceHandler = (*AuthHandler)(nil)

type AuthHandler struct {
	u usecase.Auth
}

func NewAuthHandler(u usecase.Auth) authv1connect.AuthServiceHandler {
	return &AuthHandler{u: u}
}

func (h *AuthHandler) SignUp(ctx context.Context, req *connect.Request[authv1.SignUpRequest]) (*connect.Response[authv1.SignUpResponse], error) {
	out, err := h.u.SignUp(ctx, usecase.SignUpInput{
		Email:    req.Msg.Email,
		Name:     req.Msg.Name,
		Password: req.Msg.Password,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to sign up. %w", err)
	}
	fmt.Println(out.Token)

	return connect.NewResponse(&authv1.SignUpResponse{Token: out.Token}), nil
}

func (h *AuthHandler) SignIn(ctx context.Context, req *connect.Request[authv1.SignInRequest]) (*connect.Response[authv1.SignInResponse], error) {
	out, err := h.u.SignIn(ctx, usecase.SignInInput{
		Email:    req.Msg.Email,
		Password: req.Msg.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to sing in. %w", err)
	}

	return connect.NewResponse(&authv1.SignInResponse{Token: out.Token}), nil
}
