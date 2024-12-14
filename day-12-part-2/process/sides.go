package process

import (
	"reflect"
)

type UniqueSide struct {
	neighbors []Position
	side      Direction
}

func NewUniqueSide(position Position, side Direction) *UniqueSide {
	return &UniqueSide{[]Position{position}, side}
}

func (u UniqueSide) Neighbor(position Position) bool {
	switch u.side {
	case Left, Right:
		if position.x != u.neighbors[0].x {
			return false
		}
		for _, neighbor := range u.neighbors {
			if position.y+1 == neighbor.y || position.y-1 == neighbor.y {
				return true
			}
		}
		return false
	case Up, Down:
		if position.y != u.neighbors[0].y {
			return false
		}
		for _, neighbor := range u.neighbors {
			if position.x+1 == neighbor.x || position.x-1 == neighbor.x {
				return true
			}
		}
		return false
	}
	return false
}

func (u *UniqueSide) AddNeighbor(position Position) {
	u.neighbors = append(u.neighbors, position)
}

type UniqueSides struct {
	sides []*UniqueSide
}

func internalContains(sides []*UniqueSide, direction Direction, position Position) bool {
	for _, s := range sides {
		if s.side == direction {
			if s.Neighbor(position) {
				s.AddNeighbor(position)
				return true
			}
		}
	}
	return false
}

func (u UniqueSides) Contains(direction Direction, position Position) bool {
	return internalContains(u.sides, direction, position)
}

func (u *UniqueSides) Add(side *UniqueSide) {
	u.sides = append(u.sides, side)
}

func (u *UniqueSides) GetSides() []*UniqueSide {
	return u.sides
}

func (u *UniqueSides) Count() int {
	return len(u.sides)
}

func (u *UniqueSides) Correction() bool {
	corr := make([]*UniqueSide, 0)

	for _, s := range u.sides {
		for _, n := range s.neighbors {
			if !internalContains(corr, s.side, n) {
				corr = append(corr, &UniqueSide{[]Position{n}, s.side})
			}
		}
	}

	if reflect.DeepEqual(u.sides, corr) {
		return false
	}

	u.sides = corr

	return true
}
