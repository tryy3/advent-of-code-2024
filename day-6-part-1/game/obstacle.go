package game

type Obstacle struct {
	X int
	Y int
}

func (o Obstacle) GetX() int {
	return o.X
}

func (o Obstacle) GetY() int {
	return o.Y
}

func (o Obstacle) IsObstacle() bool {
	return true
}

func (o Obstacle) GetSymbol() string {
	return "#"
}
