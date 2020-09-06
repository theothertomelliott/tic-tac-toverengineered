package spaceinmemory

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/space"
)

var _ space.Space = &Space{}

func New() space.Space {
	return &Space{}
}

func NewWithMark(m *player.Mark) space.Space {
	return &Space{
		mark: m,
	}
}

// Space implements the Space interface in memory
type Space struct {
	mark *player.Mark
}

// Mark returns the mark applied to this space, or nil if there is none
func (s *Space) Mark() (*player.Mark, error) {
	return s.mark, nil
}

// SetMark applies the specified mark to this space.
func (s *Space) SetMark(m player.Mark) error {
	s.mark = &m
	return nil
}
