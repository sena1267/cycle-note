package handler

import (
	"context"

	"connectrpc.com/connect"
	pingv1 "github.com/sena1267/cycle-note/gen/protobuf/ping/v1"
)

type PingHandler struct{}

func (h *PingHandler) Ping(ctx context.Context, req *connect.Request[pingv1.PingRequest]) (*connect.Response[pingv1.PingResponse], error) {
	return connect.NewResponse(&pingv1.PingResponse{Message: "Success"}), nil
}
