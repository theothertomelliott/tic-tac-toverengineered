package otelmonitoring

import (
	"context"
	"log"
)

// TODO: Propagate fields appropriate

func (m *Monitoring) AddFieldToSpan(ctx context.Context, key string, value interface{}) {
	log.Println("Adding field to span: ", key, value)
}

func (m *Monitoring) AddFieldToTrace(ctx context.Context, key string, value interface{}) {
	log.Println("Adding field to trace: ", key, value)
}
