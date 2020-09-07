package player

// Mark is a value representing a player's mark. Either X or O.
type Mark string

const (
	X = Mark("X")
	O = Mark("O")
)

func (m *Mark) String() string {
	if m == nil {
		return "nil"
	}
	return string(*m)
}

func MarkToPointer(m Mark) *Mark {
	return &m
}
