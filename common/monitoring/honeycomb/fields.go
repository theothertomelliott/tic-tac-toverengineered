package honeycomb

import (
	"context"

	"github.com/honeycombio/beeline-go"
)

func (m *Monitoring) AddField(ctx context.Context, key string, value interface{}) {
	beeline.AddField(ctx, key, value)
}

func (m *Monitoring) AddFieldToTrace(ctx context.Context, key string, value interface{}) {
	beeline.AddFieldToTrace(ctx, key, value)
}
