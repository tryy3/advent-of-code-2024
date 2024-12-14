package process

type Position struct {
	x int
	y int
}

func NewPosition(x, y int) *Position {
	return &Position{x, y}
}
