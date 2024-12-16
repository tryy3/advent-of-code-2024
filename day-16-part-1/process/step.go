package process

type Step struct {
	possibleMovesLeft []Direction
	position          *Position
	lastDirection     Direction
}

func (s *Step) PopDirection() Direction {
	direction := s.possibleMovesLeft[0]
	s.lastDirection = direction
	s.possibleMovesLeft = s.possibleMovesLeft[1:]
	return direction
}

func NewStep(position *Position, previousDirection Direction) *Step {
	var possibleMovesLeft []Direction
	// Prefer moving in the same direction as the previous move
	// also remove the direction that was just used
	if previousDirection == Down {
		possibleMovesLeft = []Direction{Down, Left, Right}
	} else if previousDirection == Up {
		possibleMovesLeft = []Direction{Up, Left, Right}
	} else if previousDirection == Left {
		possibleMovesLeft = []Direction{Left, Down, Up}
	} else if previousDirection == Right {
		possibleMovesLeft = []Direction{Right, Down, Up}
	}
	// possibleMovesLeft = []Direction{Down, Left, Right, Up}

	return &Step{
		possibleMovesLeft: possibleMovesLeft,
		position:          position,
	}
}
