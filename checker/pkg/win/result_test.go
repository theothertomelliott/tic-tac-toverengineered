package win_test

import (
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
)

func TestFinished(t *testing.T) {
	var tests = []struct {
		name     string
		result   win.Result
		expected bool
	}{
		{
			name:     "incomplete game",
			result:   win.Result{},
			expected: false,
		},
		{
			name: "has winner",
			result: win.Result{
				Winner: player.MarkToPointer(player.O),
			},
			expected: true,
		},
		{
			name: "draw",
			result: win.Result{
				IsDraw: true,
			},
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.result.Finished()
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
