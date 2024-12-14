package process

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

type Grid struct {
	value string
	grid  map[Position]string
}

func NewGrid(value string) *Grid {
	return &Grid{value, make(map[Position]string)}
}

func (g *Grid) Get(x, y int) string {
	return g.grid[Position{x, y}]
}

func (g *Grid) Set(x, y int, value string) {
	g.grid[Position{x, y}] = value
}

func (g *Grid) Value() string {
	return g.value
}

func (g *Grid) CountFences() int {
	fences := 0
	for pos := range g.grid {
		// check left
		leftPos := Position{pos.x - 1, pos.y}
		if _, ok := g.grid[leftPos]; !ok {
			fences++
		}
		// check right
		rightPos := Position{pos.x + 1, pos.y}
		if _, ok := g.grid[rightPos]; !ok {
			fences++
		}
		// check up
		upPos := Position{pos.x, pos.y - 1}
		if _, ok := g.grid[upPos]; !ok {
			fences++
		}
		// check down
		downPos := Position{pos.x, pos.y + 1}
		if _, ok := g.grid[downPos]; !ok {
			fences++
		}
	}

	return fences
}

func (g *Grid) FirstPosition() Position {
	next, _ := iter.Pull(maps.Keys(g.grid))
	pos, _ := next()
	return pos
}

func (g *Grid) CountSides() int {

	var traverse func(pos Position, dir Direction)
	uniqueSides := UniqueSides{}
	checked := map[Position][]Direction{}

	traverse = func(pos Position, dir Direction) {
		if _, ok2 := g.grid[pos]; ok2 {
			return
		}

		if _, ok := checked[pos]; ok {
			if slices.Contains(checked[pos], dir) {
				return
			} else {
				checked[pos] = append(checked[pos], dir)
			}
		} else {
			checked[pos] = []Direction{dir}
		}

		if !uniqueSides.Contains(dir, pos) {
			uniqueSides.Add(NewUniqueSide(pos, dir))
		}

		// fmt.Printf("x, y: %d, %d, dir: %s\n", pos.x, pos.y, dir)

		if dir == Right {
			if _, ok := g.grid[Position{pos.x - 1, pos.y + 1}]; ok {
				traverse(Position{pos.x, pos.y + 1}, Right)
			}
			if _, ok := g.grid[Position{pos.x - 1, pos.y - 1}]; ok {
				traverse(Position{pos.x, pos.y - 1}, Right)
			}
		} else if dir == Left {
			if _, ok := g.grid[Position{pos.x + 1, pos.y + 1}]; ok {
				traverse(Position{pos.x, pos.y + 1}, Left)
			}
			if _, ok := g.grid[Position{pos.x + 1, pos.y - 1}]; ok {
				traverse(Position{pos.x, pos.y - 1}, Left)
			}
		} else if dir == Up {
			if _, ok := g.grid[Position{pos.x + 1, pos.y + 1}]; ok {
				traverse(Position{pos.x + 1, pos.y}, Up)
			}
			if _, ok := g.grid[Position{pos.x - 1, pos.y + 1}]; ok {
				traverse(Position{pos.x - 1, pos.y}, Up)
			}
		} else if dir == Down {
			if _, ok := g.grid[Position{pos.x + 1, pos.y - 1}]; ok {
				traverse(Position{pos.x + 1, pos.y}, Down)
			}
			if _, ok := g.grid[Position{pos.x - 1, pos.y - 1}]; ok {
				traverse(Position{pos.x - 1, pos.y}, Down)
			}
		}
	}

	for pos, _ := range g.grid {
		for _, side := range []Direction{Left, Right, Up, Down} {
			switch side {
			case Left:
				if _, ok := g.grid[Position{pos.x - 1, pos.y}]; !ok {
					traverse(Position{pos.x - 1, pos.y}, side)
				}
			case Right:
				if _, ok := g.grid[Position{pos.x + 1, pos.y}]; !ok {
					traverse(Position{pos.x + 1, pos.y}, side)
				}
			case Up:
				if _, ok := g.grid[Position{pos.x, pos.y - 1}]; !ok {
					traverse(Position{pos.x, pos.y - 1}, side)
				}
			case Down:
				if _, ok := g.grid[Position{pos.x, pos.y + 1}]; !ok {
					traverse(Position{pos.x, pos.y + 1}, side)
				}
			}
		}
	}

	if g.value == "F" {
		fmt.Printf("uniqueSides: %v\n", uniqueSides)
	}

	return uniqueSides.Count()
}

func (g *Grid) Cost() int {
	multiplier := len(g.grid)

	// count := g.CountFences()
	sides := g.CountSides()
	fmt.Printf("area: %d, value: %s, sides: %d, multiplier: %d, cost: %d\n", len(g.grid), g.value, sides, multiplier, multiplier*sides)

	return multiplier * sides
}
