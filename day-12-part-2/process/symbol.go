package process

type Symbol struct {
	value   string
	checked bool
}

func NewSymbol(value string) *Symbol {
	return &Symbol{value, false}
}

func (s *Symbol) Check() {
	s.checked = true
}

func (s *Symbol) IsUnchecked() bool {
	return !s.checked
}

func (s *Symbol) IsChecked() bool {
	return s.checked
}

func (s *Symbol) Value() string {
	return s.value
}
