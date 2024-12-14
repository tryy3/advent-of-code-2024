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
	Paused
	Ended
)

type Game struct {
	Map    [][]Entity
	Guards []*Guard
	State  GameState
}

func NewGame() *Game {
	return &Game{
		Map:    make([][]Entity, 0),
		Guards: make([]*Guard, 0),
		State:  Running,
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
				game.Map = append(game.Map, []Entity{})
			}
			switch r {
			case '.':
				game.Map[y] = append(game.Map[y], &Empty{X: x, Y: y})
			case '#':
				game.Map[y] = append(game.Map[y], &Obstacle{X: x, Y: y})
			case '^':
				game.Map[y] = append(game.Map[y], &Empty{X: x, Y: y, StepsCount: 1})
				game.Guards = append(game.Guards, &Guard{X: x, Y: y, Direction: Up})
			case '<':
				game.Map[y] = append(game.Map[y], &Empty{X: x, Y: y, StepsCount: 1})
				game.Guards = append(game.Guards, &Guard{X: x, Y: y, Direction: Left})
			case '>':
				game.Map[y] = append(game.Map[y], &Empty{X: x, Y: y, StepsCount: 1})
				game.Guards = append(game.Guards, &Guard{X: x, Y: y, Direction: Right})
			case 'v':
				game.Map[y] = append(game.Map[y], &Empty{X: x, Y: y, StepsCount: 1})
				game.Guards = append(game.Guards, &Guard{X: x, Y: y, Direction: Down})
			}

			x++
		}
	}

	return game, nil
}

func (g *Game) Render() string {

	var output strings.Builder

	for _, row := range g.Map {
		for _, entity := range row {
			isGuard := false
			for _, guard := range g.Guards {
				if guard.X == entity.GetX() && guard.Y == entity.GetY() {
					output.WriteString(guard.GetSymbol())
					isGuard = true
					break
				}
			}
			if !isGuard {
				output.WriteString(entity.GetSymbol())
			}
		}
		output.WriteRune('\n')
	}
	if g.State == Ended {
		total := 0
		for _, row := range g.Map {
			for _, entity := range row {
				if empty, ok := entity.(*Empty); ok {
					if empty.StepsCount > 0 {
						total++
					}
				}
			}
		}
		output.WriteString(fmt.Sprintf("Game Ended. Total steps: %d", total))
	} else if g.State == Paused {
		output.WriteString("Game is paused...")
	} else {
		output.WriteString("Game is running...")
	}

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

	for _, guard := range g.Guards {
		switch guard.Direction {
		case Up:
			newY := guard.Y - 1
			if newY < 0 {
				guard.MoveUp()
				g.EndGame()
			} else if g.Map[newY][guard.X].IsObstacle() {
				guard.RotateRight()
			} else {
				guard.MoveUp()
				empty := g.Map[guard.Y][guard.X].(*Empty)
				empty.IncreaseSteps()
			}
		case Down:
			newY := guard.Y + 1
			if newY >= len(g.Map) {
				guard.MoveDown()
				g.EndGame()
			} else if g.Map[newY][guard.X].IsObstacle() {
				guard.RotateRight()
			} else {
				guard.MoveDown()
				empty := g.Map[guard.Y][guard.X].(*Empty)
				empty.IncreaseSteps()
			}
		case Left:
			newX := guard.X - 1
			if newX < 0 {
				guard.MoveLeft()
				g.EndGame()
			} else if g.Map[guard.Y][newX].IsObstacle() {
				guard.RotateRight()
			} else {
				guard.MoveLeft()
				empty := g.Map[guard.Y][guard.X].(*Empty)
				empty.IncreaseSteps()
			}
		case Right:
			newX := guard.X + 1
			if newX >= len(g.Map[guard.Y]) {
				guard.MoveRight()
				g.EndGame()
			} else if g.Map[guard.Y][newX].IsObstacle() {
				guard.RotateRight()
			} else {
				guard.MoveRight()
				empty := g.Map[guard.Y][guard.X].(*Empty)
				empty.IncreaseSteps()
			}
		}
	}
}

func (g *Game) EndGame() {
	// total := 0
	// for _, row := range g.Map {
	// 	for _, entity := range row {
	// 		if empty, ok := entity.(*Empty); ok {
	// 			if empty.StepsCount > 0 {
	// 				total++
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Printf("Game Ended. Total steps: %d\n", total)
	g.State = Ended
}
