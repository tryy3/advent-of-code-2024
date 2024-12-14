package process

type Velocity struct {
	x int
	y int
}

func NewVelocity(x, y int) *Velocity {
	return &Velocity{x, y}
}

func (v *Velocity) GetX() int {
	return v.x
}

func (v *Velocity) GetY() int {
	return v.y
}
