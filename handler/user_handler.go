package handler

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/sena1267/cycle-note/domain/model"
	userv1 "github.com/sena1267/cycle-note/gen/protobuf/user/v1"
	"github.com/sena1267/cycle-note/gen/protobuf/user/v1/userv1connect"
	"github.com/sena1267/cycle-note/usecase"
)

type UserHandler struct {
	u usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) userv1connect.UserServiceHandler {
	return &UserHandler{u: u}
}

func (h *UserHandler) Me(ctx context.Context, req *connect.Request[userv1.MeRequest]) (*connect.Response[userv1.MeResponse], error) {
	user, err := model.GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context. %w", err)
	}

	// NOTE: Name は ctx に入っていないので出力されない
	return connect.NewResponse(&userv1.MeResponse{Id: string(user.ID), Name: user.Name}), nil
}
