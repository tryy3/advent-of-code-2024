package game

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Guard struct {
	X         int
	Y         int
	Direction Direction
}

func (g Guard) GetX() int {
	return g.X
}

func (g Guard) GetY() int {
	return g.Y
}

func (g Guard) IsObstacle() bool {
	return true
}

func (g Guard) GetDirection() Direction {
	return g.Direction
}

func (g *Guard) SetDirection(direction Direction) {
	g.Direction = direction
}

func (g Guard) GetSymbol() string {
	switch g.Direction {
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
	g.Direction = (g.Direction + 1) % 4
}

func (g *Guard) MoveUp() {
	g.Y--
}

func (g *Guard) MoveDown() {
	g.Y++
}

func (g *Guard) MoveRight() {
	g.X++
}

func (g *Guard) MoveLeft() {
	g.X--
}
