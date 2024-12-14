package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type GameState int

const (
	Running GameState = iota
	Generating
	Looping
	Paused
	Ended
)

type Game struct {
	Map         [][]*Entity
	Guard       *Guard
	State       GameState
	Generations int
	Loops       int
	Routes      []*Entity
	Updates     int

	currentRoutePosition int
	startPosition        Position
}

func NewGame() *Game {
	return &Game{
		Map:                  make([][]*Entity, 0),
		Guard:                NewGuard(NewPosition(0, 0)),
		State:                Running,
		currentRoutePosition: 0, // you can hardcode this value when debugging specific route
	}
}

func LoadGameFromFile(path string) (*Game, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	game, err := LoadGameFromReader(file)
	if err != nil {
		return nil, fmt.Errorf("error loading game from file: %v", err)
	}

	return game, nil
}

func LoadGameFromReader(reader io.Reader) (*Game, error) {
	game := NewGame()
	buffer := bufio.NewReader(reader)

	x := 0
	y := 0

	for {
		r, _, err := buffer.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		if r == '\n' {
			x = 0
			y++
		} else if r != '\r' {
			if len(game.Map) <= y {
				game.Map = append(game.Map, []*Entity{})
			}
			switch r {
			case '.':
				game.Map[y] = append(game.Map[y], NewEntity(NewPosition(x, y), SpaceEntity, nil))
			case '#':
				game.Map[y] = append(game.Map[y], NewEntity(NewPosition(x, y), ObstacleEntity, nil))
			case '^':
				game.Map[y] = append(game.Map[y], NewEntityWithStep(NewPosition(x, y), SpaceEntity, Up))
				game.Guard.SetPosition(*NewPosition(x, y))
			case '<':
				game.Map[y] = append(game.Map[y], NewEntityWithStep(NewPosition(x, y), SpaceEntity, Left))
				game.Guard.SetPosition(*NewPosition(x, y))
			case '>':
				game.Map[y] = append(game.Map[y], NewEntityWithStep(NewPosition(x, y), SpaceEntity, Right))
				game.Guard.SetPosition(*NewPosition(x, y))
			case 'v':
				game.Map[y] = append(game.Map[y], NewEntityWithStep(NewPosition(x, y), SpaceEntity, Down))
				game.Guard.SetPosition(*NewPosition(x, y))
			}

			x++
		}
	}

	game.startPosition = *game.Guard.GetPosition()

	return game, nil
}

func (g *Game) Render() string {
	var output strings.Builder
	for _, row := range g.Map {
		for _, entity := range row {
			if g.Guard.GetX() == entity.GetX() && g.Guard.GetY() == entity.GetY() {
				output.WriteString(g.Guard.GetSymbol())
			} else {
				output.WriteString(entity.GetSymbol())
			}
		}
		output.WriteRune('\n')
	}
	if g.State == Ended {
		total := 0
		for _, row := range g.Map {
			for _, entity := range row {
				if entity.GetStepsCount() > 0 {
					total++
				}
			}
		}
		output.WriteString(fmt.Sprintf("Game Ended. Total steps: %d\n", total))
	} else if g.State == Paused {
		output.WriteString("Game is paused...\n")
	} else {
		output.WriteString("Game is running...\n")
	}

	output.WriteString(fmt.Sprintf("Generations: %d\n", g.Generations))
	output.WriteString(fmt.Sprintf("Loops: %d\n", g.Loops))
	output.WriteString(fmt.Sprintf("Updates: %d\n", g.Updates))
	return output.String()
}

func (g *Game) Pause() {
	g.State = Paused
}

func (g *Game) Resume() {
	g.State = Running
}

func (g *Game) Update() {
	if g.State == Ended {
		return
	}

	g.Updates++

	switch g.Guard.GetDirection() {
	case Up:
		newY := g.Guard.GetY() - 1
		if newY < 0 {
			g.Guard.MoveUp()
			g.EndGame()
		} else if g.Map[newY][g.Guard.GetX()].IsObstacle() {
			g.Guard.RotateRight()
		} else {
			g.Guard.MoveUp()
			g.Map[g.Guard.GetY()][g.Guard.GetX()].TryAddStep(StepEvent{Position: *g.Guard.GetPosition(), Direction: Up})
		}
	case Down:
		newY := g.Guard.GetY() + 1
		if newY >= len(g.Map) {
			g.Guard.MoveDown()
			g.EndGame()
		} else if g.Map[newY][g.Guard.GetX()].IsObstacle() {
			g.Guard.RotateRight()
		} else {
			g.Guard.MoveDown()
			g.Map[g.Guard.GetY()][g.Guard.GetX()].TryAddStep(StepEvent{Position: *g.Guard.GetPosition(), Direction: Down})
		}
	case Left:
		newX := g.Guard.GetX() - 1
		if newX < 0 {
			g.Guard.MoveLeft()
			g.EndGame()
		} else if g.Map[g.Guard.GetY()][newX].IsObstacle() {
			g.Guard.RotateRight()
		} else {
			g.Guard.MoveLeft()
			g.Map[g.Guard.GetY()][g.Guard.GetX()].TryAddStep(StepEvent{Position: *g.Guard.GetPosition(), Direction: Left})
		}
	case Right:
		newX := g.Guard.GetX() + 1
		if newX >= len(g.Map[g.Guard.GetY()]) {
			g.Guard.MoveRight()
			g.EndGame()
		} else if g.Map[g.Guard.GetY()][newX].IsObstacle() {
			g.Guard.RotateRight()
		} else {
			g.Guard.MoveRight()
			g.Map[g.Guard.GetY()][g.Guard.GetX()].TryAddStep(StepEvent{Position: *g.Guard.GetPosition(), Direction: Right})
		}
	}

	if g.State == Generating {
		g.CheckForLoops()
	}
}

func (g *Game) CheckForLoops() {
	steps := g.Map[g.Guard.GetY()][g.Guard.GetX()].GetSteps()
	detectedLoops := make(map[Direction]bool)
	for _, step := range steps {
		if _, ok := detectedLoops[step.Direction]; ok {
			g.State = Looping
			g.Loops++
			g.EndGame()
			return
		}
		detectedLoops[step.Direction] = true
	}
}

func (g *Game) EndGame() {
	if g.State == Running {
		g.State = Generating
		g.SaveRoute()
	}

	if !g.NewGeneration() {
		g.State = Ended
	}
}

func (g *Game) SaveRoute() {
	var mappedRoute map[Position]bool = make(map[Position]bool)
	var route []*Entity = make([]*Entity, 0)

	for _, row := range g.Map {
		for _, entity := range row {
			if entity.GetStepsCount() > 0 {
				if _, ok := mappedRoute[*entity.GetPosition()]; !ok {
					mappedRoute[*entity.GetPosition()] = true
					route = append(route, entity)
				}
			}
		}
	}

	g.Routes = route
}

func (g *Game) NewGeneration() bool {
	// Reset state
	randomizeColors()
	g.Guard.SetPosition(g.startPosition)
	g.Guard.SetDirection(Up)
	for _, row := range g.Map {
		for _, entity := range row {
			entity.ResetSteps()
			if entity.GetType() == VirtualObstacleEntity {
				entity.ChangeEntityType(SpaceEntity)
			}
		}
	}
	g.Map[g.Guard.GetY()][g.Guard.GetX()].TryAddStep(StepEvent{Position: *g.Guard.GetPosition(), Direction: g.Guard.GetDirection()}) // Increase steps on current position
	g.State = Generating

	// Increase stats
	g.Generations++

	// Set new generation
	g.currentRoutePosition++
	if g.currentRoutePosition < len(g.Routes) {
		g.Routes[g.currentRoutePosition].ChangeEntityType(VirtualObstacleEntity)
		return true
	}
	return false
}

func (g *Game) SkipUpdates(count int) {
	for i := 0; i < count; i++ {
		g.Update()
	}
}
