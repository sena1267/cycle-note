package middleware

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/sena1267/cycle-note/pkg/util/ctxtime"
)

func NewCurrentTimeInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx = ctxtime.ContextWithTime(ctx, time.Now())
			return next(ctx, req)
		}
	}

	return interceptor
}
