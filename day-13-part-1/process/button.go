package process

type Button struct {
	addX  int
	addY  int
	price int
}

func NewButton(addX, addY, price int) *Button {
	return &Button{addX, addY, price}
}
