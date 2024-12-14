package process

import "fmt"

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

func (g *Grid) Cost() int {
	multiplier := len(g.grid)

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

	fmt.Printf("area: %d, value: %s, fences: %d, multiplier: %d, cost: %d\n", len(g.grid), g.value, fences, multiplier, multiplier*fences)

	return multiplier * fences
}
