package game

import "fmt"

type Cursor struct {
	game            *Game
	currentPosition *Value
	currentRoute    []*Route
	totalExplored   int

	goals []Position
}

func NewCursor(game *Game) *Cursor {
	return &Cursor{game: game, currentRoute: []*Route{}}
}

func (c *Cursor) GetCurrentPosition() *Position {
	if c.currentPosition == nil {
		return nil
	}
	return c.currentPosition.GetPosition()
}

func (c *Cursor) GetValue() *Value {
	return c.currentPosition
}

func (c *Cursor) MoveLeft() {
	nextMove := c.game.grid[c.GetCurrentPosition().GetY()][c.GetCurrentPosition().GetX()-1]
	if len(c.currentRoute) > 0 {
		c.currentRoute[len(c.currentRoute)-1].AddExploredDirection(Left)
	}
	c.currentRoute = append(c.currentRoute, &Route{position: *NewPosition(nextMove.GetPosition().GetX(), nextMove.GetPosition().GetY())})

	c.currentPosition = nextMove
}

func (c *Cursor) MoveRight() {
	nextMove := c.game.grid[c.GetCurrentPosition().GetY()][c.GetCurrentPosition().GetX()+1]
	if len(c.currentRoute) > 0 {
		c.currentRoute[len(c.currentRoute)-1].AddExploredDirection(Right)
	}
	c.currentRoute = append(c.currentRoute, &Route{position: *NewPosition(nextMove.GetPosition().GetX(), nextMove.GetPosition().GetY())})

	c.currentPosition = nextMove
}

func (c *Cursor) MoveUp() {
	nextMove := c.game.grid[c.GetCurrentPosition().GetY()+1][c.GetCurrentPosition().GetX()]
	if len(c.currentRoute) > 0 {
		c.currentRoute[len(c.currentRoute)-1].AddExploredDirection(Up)
	}
	c.currentRoute = append(c.currentRoute, &Route{position: *NewPosition(nextMove.GetPosition().GetX(), nextMove.GetPosition().GetY())})

	c.currentPosition = nextMove
}

func (c *Cursor) MoveDown() {
	nextMove := c.game.grid[c.GetCurrentPosition().GetY()-1][c.GetCurrentPosition().GetX()]
	if len(c.currentRoute) > 0 {
		c.currentRoute[len(c.currentRoute)-1].AddExploredDirection(Down)
	}
	c.currentRoute = append(c.currentRoute, &Route{position: *NewPosition(nextMove.GetPosition().GetX(), nextMove.GetPosition().GetY())})

	c.currentPosition = nextMove
}

func (c *Cursor) MoveBack() {
	c.currentRoute = c.currentRoute[:len(c.currentRoute)-1]
	newPosition := c.currentRoute[len(c.currentRoute)-1].GetPosition()
	c.currentPosition = c.game.grid[newPosition.GetY()][newPosition.GetX()]
}

func (c *Cursor) hasPreviouslyExplored(direction Direction) bool {
	if len(c.currentRoute) == 0 {
		return false
	}
	for _, dir := range c.currentRoute[len(c.currentRoute)-1].GetExploredDirections() {
		if dir == direction {
			return true
		}
	}
	return false
}

func (c *Cursor) canMove(x, y int, direction Direction) bool {
	if c.hasPreviouslyExplored(direction) {
		return false
	}
	possibleNextMove := c.game.GetValue(x, y)

	if possibleNextMove == nil {
		return false
	}
	diff := possibleNextMove.GetValue() - c.currentPosition.GetValue()
	if diff != 1 {
		return false
	}

	return true
}

func (c *Cursor) CanMoveLeft() bool {
	return c.canMove(c.GetCurrentPosition().GetX()-1, c.GetCurrentPosition().GetY(), Left)
}

func (c *Cursor) CanMoveRight() bool {
	return c.canMove(c.GetCurrentPosition().GetX()+1, c.GetCurrentPosition().GetY(), Right)
}

func (c *Cursor) CanMoveUp() bool {
	return c.canMove(c.GetCurrentPosition().GetX(), c.GetCurrentPosition().GetY()+1, Up)
}

func (c *Cursor) CanMoveDown() bool {
	return c.canMove(c.GetCurrentPosition().GetX(), c.GetCurrentPosition().GetY()-1, Down)
}

func (c *Cursor) CanMoveBack() bool {
	return len(c.currentRoute) > 1
}

func (c *Cursor) IsCurrentPositionGoal() bool {
	if c.currentPosition.GetValue() == 9 {
		return true
	}
	return false
}

func (c *Cursor) IncreaseGoal() {
	c.goals = append(c.goals, *NewPosition(c.GetCurrentPosition().GetX(), c.GetCurrentPosition().GetY()))
	c.game.GetValue(c.GetCurrentPosition().GetX(), c.GetCurrentPosition().GetY()).IncreaseExplored()
}

func (c *Cursor) HasPossiblePath() bool {
	return c.currentPosition != nil
}

func (c *Cursor) SetNextPossiblePath(path Position) {
	c.currentPosition = c.game.GetValue(path.GetX(), path.GetY())
	c.currentRoute = append(c.currentRoute, &Route{position: *NewPosition(path.GetX(), path.GetY())})
}

func (c *Cursor) Reset() {
	fmt.Printf("goals len: %v\n", len(c.goals))
	c.currentPosition = nil
	c.currentRoute = []*Route{}
	c.goals = []Position{}
}

func (c *Cursor) GetCurrentRoute() []*Route {
	return c.currentRoute
}
