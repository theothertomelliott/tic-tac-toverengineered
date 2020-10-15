package monitoring

import (
	"context"

	"github.com/honeycombio/beeline-go"
)

func AddField(ctx context.Context, key string, value interface{}) {
	beeline.AddField(ctx, key, value)
}

func AddFieldToTrace(ctx context.Context, key string, value interface{}) {
	beeline.AddFieldToTrace(ctx, key, value)
}
