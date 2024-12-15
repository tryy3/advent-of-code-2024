package process

type Food struct {
	position *Position
}

func NewFood(position *Position) *Food {
	return &Food{position: position}
}

func (f *Food) GetPosition() *Position {
	return f.position
}

func (f *Food) SetPosition(position *Position) {
	f.position = position
}

func (f *Food) GetSymbol() string {
	return "O"
}

func (f *Food) IsWall() bool {
	return false
}

func (f *Food) IsFood() bool {
	return true
}
