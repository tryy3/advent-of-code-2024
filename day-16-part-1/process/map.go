package process

import "fmt"

type Map struct {
	start *Position
	goal  *Position

	possiblePositions map[string]*Position
	maxX              int
	maxY              int
}

func NewMap() *Map {
	return &Map{
		possiblePositions: make(map[string]*Position),
	}
}

func (m *Map) SetStart(start *Position) {
	m.start = start
}

func (m *Map) SetGoal(goal *Position) {
	m.goal = goal
}

func (m *Map) GetGoal() *Position {
	return m.goal
}

func (m *Map) GetStart() *Position {
	return m.start
}

func (m *Map) AddPosition(x, y int) {
	m.possiblePositions[fmt.Sprintf("%d,%d", x, y)] = NewPosition(x, y)
}

func (m *Map) GetPosition(x, y int) *Position {
	return m.possiblePositions[fmt.Sprintf("%d,%d", x, y)]
}

func (m *Map) GetMaxX() int {
	return m.maxX
}

func (m *Map) GetMaxY() int {
	return m.maxY
}

func (m *Map) SetMaxX(maxX int) {
	m.maxX = maxX
}

func (m *Map) SetMaxY(maxY int) {
	m.maxY = maxY
}
