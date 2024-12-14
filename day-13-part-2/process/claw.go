package process

type Claw struct {
	buttonA         *Button
	buttonB         *Button
	prize           *Position
	currentPosition *Position
	totalCost       int

	totalButtonBCount int
	totalButtonACount int
}

func NewClaw(buttonA *Button, buttonB *Button, prize *Position) *Claw {
	return &Claw{buttonA, buttonB, prize, &Position{0, 0}, 0, 0, 0}
}

func (c *Claw) MoveClawForward(button *Button, steps int) {
	c.currentPosition.Add(button.addX*steps, button.addY*steps)
	c.totalCost += button.price * steps
}

func (c *Claw) MoveClawBackward(button *Button, steps int) {
	c.currentPosition.Add(-button.addX*steps, -button.addY*steps)
	c.totalCost -= button.price * steps
}

func (c *Claw) MoveClawAForward(steps int) {
	c.MoveClawForward(c.buttonA, steps)
	c.totalButtonACount += steps
}

func (c *Claw) MoveClawBForward(steps int) {
	c.MoveClawForward(c.buttonB, steps)
	c.totalButtonBCount += steps
}

func (c *Claw) MoveClawABackward(steps int) {
	c.MoveClawBackward(c.buttonA, steps)
	c.totalButtonACount -= steps
}

func (c *Claw) MoveClawBBackward(steps int) {
	c.MoveClawBackward(c.buttonB, steps)
	c.totalButtonBCount -= steps
}

func (c *Claw) GetCurrentPosition() *Position {
	return c.currentPosition
}

func (c *Claw) GetTotalCost() int {
	return c.totalCost
}

func (c *Claw) GetPrizePosition() *Position {
	return c.prize
}

func (c *Claw) IsPrizeReached() bool {
	return c.currentPosition.GetX() == c.prize.GetX() && c.currentPosition.GetY() == c.prize.GetY()
}

func (c *Claw) IsPositionFurtherThanPrize() bool {
	return c.currentPosition.GetX() >= c.prize.GetX() || c.currentPosition.GetY() >= c.prize.GetY()
}

func (c *Claw) IsPositionBeforePrize() bool {
	return c.currentPosition.GetX() <= c.prize.GetX() && c.currentPosition.GetY() <= c.prize.GetY()
}

func (c *Claw) GetTotalButtonBCount() int {
	return c.totalButtonBCount
}

func (c *Claw) GetTotalButtonACount() int {
	return c.totalButtonACount
}
