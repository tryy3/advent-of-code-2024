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

func (c *Claw) MoveClawForward(button *Button) {
	c.currentPosition.Add(button.addX, button.addY)
	c.totalCost += button.price
}

func (c *Claw) MoveClawBackward(button *Button) {
	c.currentPosition.Add(-button.addX, -button.addY)
	c.totalCost -= button.price
}

func (c *Claw) MoveClawAForward() {
	c.MoveClawForward(c.buttonA)
	c.totalButtonACount++
}

func (c *Claw) MoveClawBForward() {
	c.MoveClawForward(c.buttonB)
	c.totalButtonBCount++
}

func (c *Claw) MoveClawABackward() {
	c.MoveClawBackward(c.buttonA)
	c.totalButtonACount--
}

func (c *Claw) MoveClawBBackward() {
	c.MoveClawBackward(c.buttonB)
	c.totalButtonBCount--
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
