package game

type Position struct {
	x int
	y int
}

func (p Position) GetX() int {
	return p.x
}

func (p Position) GetY() int {
	return p.y
}

func (p *Position) SetX(x int) {
	p.x = x
}

func (p *Position) SetY(y int) {
	p.y = y
}

func NewPosition(x int, y int) *Position {
	return &Position{x: x, y: y}
}
