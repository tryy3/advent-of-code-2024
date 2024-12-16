package process

type Start struct {
	position *Position
}

func NewStart(position *Position) *Start {
	return &Start{position: position}
}
