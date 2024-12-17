package process

type Register struct {
	value int
}

func NewRegister(value int) *Register {
	return &Register{value: value}
}

func (r *Register) SetValue(value int) {
	r.value = value
}

func (r *Register) GetValue() int {
	return r.value
}
