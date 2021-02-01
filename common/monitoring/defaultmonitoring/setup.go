package defaultmonitoring

import "github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/honeycomb"

func Init(serviceName string) {
	honeycomb.Init(serviceName)
}
