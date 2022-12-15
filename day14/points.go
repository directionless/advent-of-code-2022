package day14

type square int

const (
	emptySquare square = iota
	rockSquare
	sandSquare
	sourceSquare
	airSquare
)

func (s square) String() string {
	switch s {
	case emptySquare:
		return " "
	case rockSquare:
		return "#"
	case sandSquare:
		return "o"
	case sourceSquare:
		return "+"
	case airSquare:
		return "."
	default:
		return " "
	}
}

var (
	Rock       = square(rockSquare)
	Sand       = square(sandSquare)
	SandSource = square(sourceSquare)
	Air        = square(airSquare)
)
