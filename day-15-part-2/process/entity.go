package process

type Entity interface {
	GetPosition() *Position
	SetPosition(position *Position)
	GetSymbol() string
	IsWall() bool
	IsFood() bool
}
