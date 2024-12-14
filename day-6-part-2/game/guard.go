package game

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func NewGuard(position *Position) *Guard {
	return &Guard{position: position}
}

type Guard struct {
	position  *Position
	direction Direction
}

func (g Guard) GetX() int {
	return g.position.GetX()
}

func (g Guard) GetY() int {
	return g.position.GetY()
}

func (g Guard) GetPosition() *Position {
	return g.position
}

func (g Guard) GetDirection() Direction {
	return g.direction
}

func (g *Guard) SetPosition(position Position) {
	g.position.SetX(position.GetX())
	g.position.SetY(position.GetY())
}

func (g *Guard) SetDirection(direction Direction) {
	g.direction = direction
}

func (g Guard) GetSymbol() string {
	switch g.direction {
	case Up:
		return "^"
	case Right:
		return ">"
	case Down:
		return "v"
	case Left:
		return "<"
	}
	return " "
}

func (g *Guard) RotateRight() {
	g.direction = (g.direction + 1) % 4
}

func (g *Guard) MoveUp() {
	g.position.y--
}

func (g *Guard) MoveDown() {
	g.position.y++
}

func (g *Guard) MoveRight() {
	g.position.x++
}

func (g *Guard) MoveLeft() {
	g.position.x--
}
