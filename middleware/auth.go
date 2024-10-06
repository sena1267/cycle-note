package middleware

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"connectrpc.com/connect"
	"github.com/sena1267/cycle-note/domain/model"
	"github.com/sena1267/cycle-note/pkg/auth"
)

const tokenHeader = "Authorization"

func NewAuthInterceptor(authenticator auth.Authenticator) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if slices.Contains([]string{"SignUp", "SignIn"}, strings.Split(req.Spec().Procedure, "/")[2]) {
				return next(ctx, req)
			}

			if req.Header().Get(tokenHeader) == "" {
				// Check token in handlers.
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}

			authHeader := strings.Split(req.Header().Get("Authorization"), " ")
			if len(authHeader) != 2 {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("Authorization header should have value: 'Bearer <token>' "),
				)
			}

			token := authHeader[1]

			userID, err := authenticator.GetUserIDFromToken(token)
			if err != nil {
				return nil, fmt.Errorf("failed to get user id from token. %w", err)
			}
			fmt.Println(userID)

			return next(model.ContextWithUser(ctx, model.User{ID: userID}), req)
		}
	}
	return interceptor
}
