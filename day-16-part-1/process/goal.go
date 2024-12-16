package process

type Goal struct {
	position *Position
}

func NewGoal(position *Position) *Goal {
	return &Goal{position: position}
}

func (g *Goal) GetPosition() *Position {
	return g.position
}
