package game

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

type Route struct {
	position Position

	exploredDirections []Direction
}

func (r *Route) HasExploredDirection(direction Direction) bool {
	for _, explored := range r.exploredDirections {
		if explored == direction {
			return true
		}
	}
	return false
}

func (r *Route) AddExploredDirection(direction Direction) {
	r.exploredDirections = append(r.exploredDirections, direction)
}

func (r *Route) GetExploredDirections() []Direction {
	return r.exploredDirections
}

func (r *Route) GetPosition() Position {
	return r.position
}
