package monitoring

import (
	"context"
)

func AddField(ctx context.Context, key string, value interface{}) {
	Default.AddField(ctx, key, value)
}

func AddFieldToTrace(ctx context.Context, key string, value interface{}) {
	Default.AddFieldToTrace(ctx, key, value)
}
