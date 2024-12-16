package process

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

type Processor struct {
	grid          *Map
	path          []Step
	recordedMoves [][]Direction
	isOver        bool
	currentCost   int
	lowestCost    int
	steppedOn     map[string]int
	priceOfSteps  map[string]int
	maxSteppedOn  int
	positions     map[string]bool
}

func NewProcessor() *Processor {
	return &Processor{
		grid:          NewMap(),
		path:          []Step{},
		recordedMoves: [][]Direction{},
		isOver:        false,
		currentCost:   0,
		lowestCost:    200000, // 194704
		steppedOn:     map[string]int{},
		priceOfSteps:  map[string]int{},
		maxSteppedOn:  5000,
		positions:     map[string]bool{},
	}
}

func (p *Processor) SetLowestCost(cost int) {
	p.lowestCost = cost
}

func (p *Processor) SetMaxSteppedOn(max int) {
	p.maxSteppedOn = max
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
	r := bufio.NewReader(reader)

	var x, y int
	for {
		rune, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				y++
				if y > p.grid.GetMaxY() {
					p.grid.SetMaxY(y)
				}
				break
			}
			return fmt.Errorf("failed to read rune: %w", err)
		}

		if rune == '\n' {
			x = 0
			y++
			if y > p.grid.GetMaxY() {
				p.grid.SetMaxY(y)
			}
			continue
		}
		if rune == '\r' {
			continue
		}

		if rune == '.' {
			p.grid.AddPosition(x, y)
		}
		if rune == 'S' {
			p.grid.AddPosition(x, y)
			p.grid.SetStart(p.grid.GetPosition(x, y))
		}
		if rune == 'E' {
			p.grid.AddPosition(x, y)
			p.grid.SetGoal(p.grid.GetPosition(x, y))
		}
		x++
		if x > p.grid.GetMaxX() {
			p.grid.SetMaxX(x)
		}
	}

	return nil
}

func (p *Processor) IsOver() bool {
	return p.isOver
}

func (p *Processor) Begin() {
	p.path = append(p.path, *NewStep(p.grid.GetStart(), Right))
}

func (p *Processor) Update() {
	if len(p.path) <= 0 {
		p.isOver = true
		return
	}
	if !p.MoveForward() {
		p.MoveBack()
	}

	p.currentCost = p.CurrentPathCost()
	// fmt.Printf("Current cost: %d, Lowest cost: %d\n", p.currentCost, p.lowestCost)
	if p.currentCost > p.lowestCost {
		p.MoveBack()
		return
	}

	// p.Render()
}

func (p *Processor) MoveBack() {
	p.path = p.path[:len(p.path)-1]
}

func (p *Processor) CurrentPathCost() int {
	moves := []Direction{}
	lastDirection := Right
	for _, s := range p.path {
		if s.lastDirection != lastDirection {
			moves = append(moves, Rotate)
			lastDirection = s.lastDirection
		}
		moves = append(moves, s.lastDirection)

	}
	return SumPath(moves)
}

func (p *Processor) CopyPath() {
	moves := []Direction{}
	lastDirection := Right
	for _, s := range p.path {
		if s.lastDirection != lastDirection {
			moves = append(moves, Rotate)
			lastDirection = s.lastDirection
		}
		moves = append(moves, s.lastDirection)
		p.positions[fmt.Sprintf("%d,%d", s.position.GetX(), s.position.GetY())] = true
	}
	p.recordedMoves = append(p.recordedMoves, moves)
	// sum := SumPath(moves)
	// fmt.Println(sum)
}

func (p *Processor) MoveForward() bool {
	for {
		if len(p.path[len(p.path)-1].possibleMovesLeft) <= 0 {
			return false
		}
		nextPossibleMove := p.path[len(p.path)-1].PopDirection()
		currentPosition := p.path[len(p.path)-1].position

		var newPosition *Position
		switch nextPossibleMove {
		case Down:
			newPosition = p.grid.GetPosition(currentPosition.GetX(), currentPosition.GetY()+1)
		case Up:
			newPosition = p.grid.GetPosition(currentPosition.GetX(), currentPosition.GetY()-1)
		case Left:
			newPosition = p.grid.GetPosition(currentPosition.GetX()-1, currentPosition.GetY())
		case Right:
			newPosition = p.grid.GetPosition(currentPosition.GetX()+1, currentPosition.GetY())
		}

		if newPosition == nil {
			continue
		}

		// fmt.Printf("new Position: %d,%d - Goal: %d,%d\n", newPosition.GetX(), newPosition.GetY(), p.grid.GetGoal().GetX(), p.grid.GetGoal().GetY())
		if newPosition.GetX() == p.grid.GetGoal().GetX() && newPosition.GetY() == p.grid.GetGoal().GetY() {
			p.CopyPath()
			// p.Render()
			newLowestCost := p.CurrentPathCost()
			if newLowestCost <= p.lowestCost {
				p.lowestCost = newLowestCost
				fmt.Printf("New lowest cost: %d\n", p.lowestCost)
			}
			return false
		}

		visited := false
		for _, s := range p.path {
			if s.position.GetX() == newPosition.GetX() && s.position.GetY() == newPosition.GetY() {
				visited = true
				break
			}
		}
		if visited {
			continue
		}

		// p.currentCost = p.CurrentPathCost()
		p.path = append(p.path, *NewStep(newPosition, nextPossibleMove))
		p.StepOn(newPosition)

		// if p.steppedOn[fmt.Sprintf("%d,%d", newPosition.GetX(), newPosition.GetY())] > p.maxSteppedOn {
		// 	return false
		// }
		// fmt.Printf("Current cost: %d, Price of steps: %d\n", p.currentCost, p.priceOfSteps[fmt.Sprintf("%d,%d", newPosition.GetX(), newPosition.GetY())])

		// if p.currentCost > p.priceOfSteps[fmt.Sprintf("%d,%d", newPosition.GetX(), newPosition.GetY())] {
		// 	return false
		// }
		return true
	}
}

func (p *Processor) StepOn(position *Position) {
	key := fmt.Sprintf("%d,%d", position.GetX(), position.GetY())
	if _, ok := p.steppedOn[key]; !ok {
		p.steppedOn[key] = 0
	}
	p.steppedOn[key]++

	if _, ok := p.priceOfSteps[key]; !ok {
		p.priceOfSteps[key] = math.MaxInt
	}
	if p.priceOfSteps[key] > p.currentCost {
		p.priceOfSteps[key] = p.currentCost
	}
}

func (p *Processor) Render() {
	start := p.grid.GetStart()
	goal := p.grid.GetGoal()
	for y := 0; y < p.grid.GetMaxY(); y++ {
		for x := 0; x < p.grid.GetMaxX(); x++ {
			pos := p.grid.GetPosition(x, y)
			if start.GetX() == x && start.GetY() == y {
				fmt.Print("S")
			} else if goal.GetX() == x && goal.GetY() == y {
				fmt.Print("E")
			} else if pos != nil {
				found := false
				for _, s := range p.path {
					if s.position.GetX() == x && s.position.GetY() == y {
						found = true
						switch s.lastDirection {
						case Down:
							fmt.Print("v")
						case Up:
							fmt.Print("^")
						case Left:
							fmt.Print("<")
						case Right:
							fmt.Print(">")
						}
						break
					}
				}
				if !found {
					fmt.Print(".")
				}
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func SumPath(moves []Direction) int {
	sum := 0
	for _, move := range moves {
		if move == Rotate {
			sum += 1000
		} else {
			sum++
		}
	}
	return sum
}

func (p *Processor) SumPaths() {
	lowestSum := math.MaxInt
	for _, moves := range p.recordedMoves {
		sum := 0
		for _, move := range moves {
			if move == Rotate {
				sum += 1000
			} else {
				sum++
			}
		}
		if sum < lowestSum {
			lowestSum = sum
		}
	}
	fmt.Println(lowestSum)
	fmt.Println(len(p.recordedMoves))
	fmt.Println(len(p.positions))
}
