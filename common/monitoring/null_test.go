package monitoring_test

import (
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
)

func TestNullClose(t *testing.T) {
	err := monitoring.Close()
	if err != nil {
		t.Error(err)
	}
}
