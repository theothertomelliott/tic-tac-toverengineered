package monitoring

import (
	"context"
)

func AddFieldToSpan(ctx context.Context, key string, value interface{}) {
	Default.AddFieldToSpan(ctx, key, value)
}

func AddFieldToTrace(ctx context.Context, key string, value interface{}) {
	Default.AddFieldToTrace(ctx, key, value)
}
