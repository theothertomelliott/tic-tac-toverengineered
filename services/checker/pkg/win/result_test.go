package win_test

import (
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win"
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

func TestEqual(t *testing.T) {
	var tests = []struct {
		name     string
		resultA  win.Result
		resultB  win.Result
		expected bool
	}{
		{
			name:     "incomplete game",
			resultA:  win.Result{},
			resultB:  win.Result{},
			expected: true,
		},
		{
			name: "same winner",
			resultA: win.Result{
				Winner: player.MarkToPointer(player.O),
			},
			resultB: win.Result{
				Winner: player.MarkToPointer(player.O),
			},
			expected: true,
		},
		{
			name: "different winner",
			resultA: win.Result{
				Winner: player.MarkToPointer(player.O),
			},
			resultB: win.Result{
				Winner: player.MarkToPointer(player.X),
			},
			expected: false,
		},
		{
			name:    "one with winner, one without winner",
			resultA: win.Result{},
			resultB: win.Result{
				Winner: player.MarkToPointer(player.X),
			},
			expected: false,
		},
		{
			name: "both draw",
			resultA: win.Result{
				IsDraw: true,
			},
			resultB: win.Result{
				IsDraw: true,
			},
			expected: true,
		},
		{
			name: "one draw, one not",
			resultA: win.Result{
				IsDraw: true,
			},
			resultB: win.Result{
				IsDraw: false,
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.resultA.Equal(test.resultB)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
