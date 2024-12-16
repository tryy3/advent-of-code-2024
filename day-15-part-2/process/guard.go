package process

type Guard struct {
	position *Position
	size     *Size
}

func NewGuard(position *Position) *Guard {
	return &Guard{position: position, size: NewSize(1, 1)}
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

func (g *Guard) GetSize() *Size {
	return g.size
}

func (g *Guard) Move(direction Direction) {
	switch direction {
	case Up:
		g.position.y -= 1
	case Down:
		g.position.y += 1
	case Left:
		g.position.x -= 1
	case Right:
		g.position.x += 1
	}
}
