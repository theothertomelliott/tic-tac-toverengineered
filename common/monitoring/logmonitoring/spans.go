package logmonitoring

import (
	"context"
	"log"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
)

func (m *Monitoring) StartSpan(ctx context.Context, name string) (context.Context, monitoring.Span) {
	log.Println("Starting span: ", name)
	return ctx, &span{name: name}
}

type span struct {
	name string
}

func (s *span) Finish() error {
	log.Println("Finished span: ", s.name)
	return nil
}
