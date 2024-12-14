package process

import "fmt"

type Position struct {
	x int
	y int
}

func (p *Position) GetY() int {
	return p.y
}

func (p *Position) GetX() int {
	return p.x
}

func (p *Position) Add(x, y int) {
	p.x += x
	p.y += y
}

func NewPosition(x, y int) *Position {
	return &Position{x, y}
}

func (p *Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}
