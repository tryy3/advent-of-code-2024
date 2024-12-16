package process

type Wall struct {
	position *Position
	size     *Size
}

func NewWall(position *Position) *Wall {
	return &Wall{position: position, size: NewSize(2, 1)}
}

func (w *Wall) GetPosition() *Position {
	return w.position
}

func (w *Wall) SetPosition(position *Position) {
	w.position = position
}

func (w *Wall) GetSymbol() string {
	return "##"
}

func (w *Wall) IsWall() bool {
	return true
}

func (w *Wall) IsFood() bool {
	return false
}

func (w *Wall) GetSize() *Size {
	return w.size
}

func (w *Wall) Move(direction Direction) {
	switch direction {
	case Up:
		w.position.y -= 1
	case Down:
		w.position.y += 1
	case Left:
		w.position.x -= 1
	case Right:
		w.position.x += 1
	}
}
