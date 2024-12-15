package process

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) String() string {
	return []string{"^", ">", "v", "<"}[d]
}
