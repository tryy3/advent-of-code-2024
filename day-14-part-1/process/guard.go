package process

type Guard struct {
	position *Position
	velocity *Velocity
}

func NewGuard(position *Position, velocity *Velocity) *Guard {
	return &Guard{position, velocity}
}

func (g *Guard) GetPosition() *Position {
	return g.position
}

func (g *Guard) GetVelocity() *Velocity {
	return g.velocity
}
