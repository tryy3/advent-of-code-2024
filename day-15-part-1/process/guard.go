package process

type Guard struct {
	position *Position
}

func NewGuard(position *Position) *Guard {
	return &Guard{position: position}
}

func (g *Guard) GetSymbol() string {
	return "@"
}

func (g *Guard) GetPosition() *Position {
	return g.position
}

func (g *Guard) SetPosition(position *Position) {
	g.position = position
}

func (g *Guard) IsWall() bool {
	return false
}

func (g *Guard) IsFood() bool {
	return false
}
