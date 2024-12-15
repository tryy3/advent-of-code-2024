package process

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Processor struct {
	maxWidth         int
	maxHeight        int
	entities         []Entity
	guard            *Guard
	directions       []Direction
	currentDirection int
}

func NewProcessor() *Processor {
	return &Processor{
		maxWidth:         0,
		maxHeight:        0,
		entities:         []Entity{},
		directions:       []Direction{},
		currentDirection: 0,
	}
}

func (p *Processor) LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return p.LoadFromReader(file)
}

func (p *Processor) LoadFromReader(reader io.Reader) error {
	scanner := bufio.NewReader(reader)

	loadMap := true

	var x, y int
	var lastRune rune
	for {
		rune, _, err := scanner.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if loadMap {
			if x > p.maxWidth {
				p.maxWidth = x
			}
			if y > p.maxHeight {
				p.maxHeight = y
			}
		}

		if rune == '\n' {
			if lastRune == '\n' {
				loadMap = false

			} else {
				x = 0
				y++
				lastRune = rune
			}
			continue
		}
		if rune == '\r' {
			continue
		}
		lastRune = rune

		if loadMap {
			switch rune {
			case '#':
				p.entities = append(p.entities, NewWall(NewPosition(x, y)))
			case 'O':
				p.entities = append(p.entities, NewFood(NewPosition(x, y)))
			case '@':
				p.guard = NewGuard(NewPosition(x, y))
				p.entities = append(p.entities, p.guard)
			}
			x++
			continue
		} else {
			switch rune {
			case '^':
				p.directions = append(p.directions, Up)
			case '>':
				p.directions = append(p.directions, Right)
			case 'v':
				p.directions = append(p.directions, Down)
			case '<':
				p.directions = append(p.directions, Left)
			}
		}

	}

	return nil
}

func (p *Processor) Update() {
	dir := p.directions[p.currentDirection]
	nextPosition := p.NextPosition(p.guard.GetPosition(), dir)
	if p.AttemptMove(p.guard.GetPosition(), dir) {
		p.guard.SetPosition(nextPosition)
	}
	p.currentDirection = (p.currentDirection + 1) % len(p.directions)
}

func (p *Processor) NextPosition(currentPosition *Position, direction Direction) *Position {
	var newPosition *Position
	switch direction {
	case Up:
		newPosition = NewPosition(currentPosition.x, currentPosition.y-1)
	case Down:
		newPosition = NewPosition(currentPosition.x, currentPosition.y+1)
	case Left:
		newPosition = NewPosition(currentPosition.x-1, currentPosition.y)
	case Right:
		newPosition = NewPosition(currentPosition.x+1, currentPosition.y)
	}
	return newPosition
}

func (p *Processor) AttemptMove(currentPosition *Position, direction Direction) bool {
	newPosition := p.NextPosition(currentPosition, direction)
	entity := p.GetEntity(newPosition.x, newPosition.y)
	if entity == nil {
		// Return true but nothing moves (no collision)
		return true
	} else {
		if entity.IsWall() {
			return false
		} else {
			if p.AttemptMove(newPosition, direction) {
				nextPosition := p.NextPosition(newPosition, direction)
				entity.SetPosition(nextPosition)
				return true
			}
		}
	}
	return false
}

func (p *Processor) Render() {
	for y := 0; y < p.maxHeight; y++ {
		for x := 0; x < p.maxWidth; x++ {
			if p.IsGuardAt(x, y) {
				fmt.Print("@")
			} else {
				entity := p.GetEntity(x, y)
				if entity != nil {
					fmt.Print(entity.GetSymbol())

				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

func (p *Processor) GetEntity(x, y int) Entity {
	for _, entity := range p.entities {
		if entity.GetPosition().x == x && entity.GetPosition().y == y {
			return entity
		}
	}
	return nil
}

func (p *Processor) IsGuardAt(x, y int) bool {
	return p.guard.GetPosition().x == x && p.guard.GetPosition().y == y
}

func (p *Processor) GetStepCount() int {
	return len(p.directions)
}

func (p *Processor) GetSum() int {
	sum := 0
	for _, entity := range p.entities {
		if entity.IsFood() {
			sum += 100*entity.GetPosition().y + entity.GetPosition().x
		}
	}
	return sum
}
