package defaultmonitoring

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/logmonitoring"
)

func Init(serviceName string) {
	//honeycomb.Init(serviceName)
	monitoring.Default = &logmonitoring.Monitoring{}
}
