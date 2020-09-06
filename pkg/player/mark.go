package player

// Mark is a value representing a player's mark. Either X or O.
type Mark int8

const (
	X Mark = iota
	O
)

func (m *Mark) String() string {
	if m == nil {
		return "nil"
	}
	if *m == X {
		return "X"
	}
	if *m == O {
		return "O"
	}
	return ""
}

func MarkToPointer(m Mark) *Mark {
	return &m
}
