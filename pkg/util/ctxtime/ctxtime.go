package ctxtime

import (
	"context"
	"time"
)

type contextTimeKey struct{}

func Now(ctx context.Context) time.Time {
	if value := ctx.Value(contextTimeKey{}); value != nil {
		if ctxTime, ok := value.(time.Time); ok {
			return ctxTime
		}
	}

	return time.Now()

}

func ContextWithTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, contextTimeKey{}, t)
}
