package process

type Size struct {
	width  int
	height int
}

func NewSize(width, height int) *Size {
	return &Size{width: width, height: height}
}

func (s *Size) GetWidth() int {
	return s.width
}

func (s *Size) GetHeight() int {
	return s.height
}
