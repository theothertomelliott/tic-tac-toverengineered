package defaultmonitoring

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/otelmonitoring"
)

func Init(serviceName string) {
	var err error
	monitoring.Default, err = otelmonitoring.New()
	if err != nil {
		panic(err)
	}
}
