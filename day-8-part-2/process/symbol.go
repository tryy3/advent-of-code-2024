package process

import "fmt"

type Symbol struct {
	value string
	x, y  int
}

func NewSymbol(value string, x, y int) *Symbol {
	return &Symbol{value: value, x: x, y: y}
}

func (s *Symbol) String() string {
	return s.value
}

func (s *Symbol) GetX() int {
	return s.x
}

func (s *Symbol) GetY() int {
	return s.y
}

func (s *Symbol) GetID() string {
	return fmt.Sprintf("x:%d;y:%d", s.GetX(), s.GetY())
}
