package logmonitoring

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
)

var _ monitoring.Monitoring = &Monitoring{}

type Monitoring struct {
}

func (m *Monitoring) Close() error {
	return nil
}
