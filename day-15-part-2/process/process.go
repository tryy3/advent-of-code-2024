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
			var entity Entity
			switch rune {
			case '#':
				entity = NewWall(NewPosition(x, y))
			case 'O':
				entity = NewFood(NewPosition(x, y))
			case '@':
				entity = NewGuard(NewPosition(x, y))
				p.guard = entity.(*Guard)
			}
			if entity != nil {
				p.entities = append(p.entities, entity)
				x += 2
			} else {
				x += 2
			}
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

	fmt.Println(p.maxWidth, p.maxHeight)

	return nil
}

func (p *Processor) Update() {
	dir := p.directions[p.currentDirection]
	// nextPosition := p.NextPosition(p.guard.GetPosition(), 1, dir)
	entities := p.AttemptMove(p.guard.GetPosition(), 1, dir)
	if entities != nil {
		for entity, _ := range entities {
			entity.Move(dir)
		}
		p.guard.Move(dir)
	}
	p.currentDirection++
}

func (p *Processor) NextPosition(currentPosition *Position, movement int, direction Direction) *Position {
	var newPosition *Position
	switch direction {
	case Up:
		newPosition = NewPosition(currentPosition.x, currentPosition.y-movement)
	case Down:
		newPosition = NewPosition(currentPosition.x, currentPosition.y+movement)
	case Left:
		newPosition = NewPosition(currentPosition.x-movement, currentPosition.y)
	case Right:
		newPosition = NewPosition(currentPosition.x+movement, currentPosition.y)
	}
	return newPosition
}

func (p *Processor) GetMovement(direction Direction, entity Entity) int {
	switch direction {
	case Up, Down:
		return entity.GetSize().GetHeight()
	case Left, Right:
		return entity.GetSize().GetWidth()
	}
	return 0
}

func (p *Processor) AttemptMove(currentPosition *Position, movement int, direction Direction) map[Entity]bool {
	newPosition := p.NextPosition(currentPosition, movement, direction)
	entity := p.GetEntity(newPosition.x, newPosition.y)
	if entity == nil {
		// Return true but nothing moves (no collision)
		return map[Entity]bool{}
	} else {
		if entity.IsWall() {
			return nil
		} else {
			entities := map[Entity]bool{entity: true}

			if direction == Up || direction == Down {
				for i := 0; i < entity.GetSize().GetWidth(); i++ {
					e := p.AttemptMove(NewPosition(entity.GetPosition().x+i, entity.GetPosition().y), 1, direction)
					if e == nil {
						return nil
					}
					for k, v := range e {
						entities[k] = v
					}
				}
			} else {
				e := p.AttemptMove(newPosition, p.GetMovement(direction, entity), direction)
				if e == nil {
					return nil
				}
				for k, v := range e {
					entities[k] = v
				}
			}
			return entities
		}
	}
	return map[Entity]bool{}
}

func (p *Processor) Render() {
	var x, y int
	for {
		if y >= p.maxHeight {
			break
		}
		for {
			if x >= p.maxWidth {
				x = 0
				break
			}
			if p.IsGuardAt(x, y) {
				fmt.Print("@")
				x++
			} else {
				entity := p.GetEntity(x, y)
				if entity != nil {
					fmt.Print(entity.GetSymbol())
					x += entity.GetSize().GetWidth()
				} else {
					fmt.Print(".")
					x++
				}
			}
		}
		y++
		fmt.Println()
	}
}

func (p *Processor) GetEntity(x, y int) Entity {
	for _, entity := range p.entities {
		entityX := entity.GetPosition().x
		maxEntityX := entityX + entity.GetSize().GetWidth()
		entityY := entity.GetPosition().y
		maxEntityY := entityY + entity.GetSize().GetHeight()

		if x >= entityX && x < maxEntityX && y >= entityY && y < maxEntityY {
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
