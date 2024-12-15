package process

type Wall struct {
	position *Position
}

func NewWall(position *Position) *Wall {
	return &Wall{position: position}
}

func (w *Wall) GetPosition() *Position {
	return w.position
}

func (w *Wall) SetPosition(position *Position) {
	w.position = position
}

func (w *Wall) GetSymbol() string {
	return "#"
}

func (w *Wall) IsWall() bool {
	return true
}

func (w *Wall) IsFood() bool {
	return false
}
