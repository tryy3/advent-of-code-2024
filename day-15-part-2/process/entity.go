package process

type Entity interface {
	GetPosition() *Position
	SetPosition(position *Position)
	GetSize() *Size
	GetSymbol() string
	IsWall() bool
	IsFood() bool
	Move(direction Direction)
}
