package process

type Food struct {
	position *Position
	size     *Size
}

func NewFood(position *Position) *Food {
	return &Food{position: position, size: NewSize(2, 1)}
}

func (f *Food) GetPosition() *Position {
	return f.position
}

func (f *Food) SetPosition(position *Position) {
	f.position = position
}

func (f *Food) GetSymbol() string {
	return "[]"
}

func (f *Food) IsWall() bool {
	return false
}

func (f *Food) IsFood() bool {
	return true
}

func (f *Food) GetSize() *Size {
	return f.size
}

func (f *Food) Move(direction Direction) {
	switch direction {
	case Up:
		f.position.y -= 1
	case Down:
		f.position.y += 1
	case Left:
		f.position.x -= 1
	case Right:
		f.position.x += 1
	}
}
