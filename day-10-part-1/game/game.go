package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Game struct {
	isOver        bool
	grid          [][]*Value
	cursor        *Cursor
	possiblePaths []Position
	checkedPaths  []Position
}

func NewGame() *Game {
	game := &Game{}
	game.cursor = NewCursor(game)
	return game
}

func LoadGameFromFile(path string) (*Game, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return LoadGameFromReader(file)
}

func LoadGameFromReader(reader io.Reader) (*Game, error) {
	game := NewGame()
	buffer := bufio.NewReader(reader)

	var x, y int
	for {
		r, _, err := buffer.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		if r == '\n' {
			y++
			x = 0
			continue
		}
		if r == '\r' {
			continue
		}

		var num int
		if r == '.' {
			num = -1
		} else {
			num, err = strconv.Atoi(string(r))
			if err != nil {
				return nil, fmt.Errorf("error converting rune to int: %v", err)
			}
		}

		val := NewValue(NewPosition(x, y), num)

		if num == 0 {
			game.possiblePaths = append(game.possiblePaths, *NewPosition(x, y))
		}

		game.AddValue(val)
		x++
	}

	return game, nil
}

func (g *Game) AddValue(val *Value) {
	y := val.GetPosition().GetY()
	if g.grid == nil {
		g.grid = make([][]*Value, 0)
	}
	if len(g.grid)-1 < y {
		g.grid = append(g.grid, make([]*Value, 0))
	}
	g.grid[y] = append(g.grid[y], val)
}

func (g *Game) IsOver() bool {
	return g.isOver
}

func (g *Game) SetNextPossiblePath() {
	g.cursor.SetNextPossiblePath(g.possiblePaths[len(g.possiblePaths)-1])

	lastPossiblePath := g.possiblePaths[len(g.possiblePaths)-1]
	g.checkedPaths = append(g.checkedPaths, lastPossiblePath)
	g.possiblePaths = g.possiblePaths[:len(g.possiblePaths)-1]
}

func (g *Game) Update() {
	if !g.cursor.HasPossiblePath() {
		g.SetNextPossiblePath()
	}

	if g.cursor.CanMoveLeft() {
		g.cursor.MoveLeft()
	} else if g.cursor.CanMoveRight() {
		g.cursor.MoveRight()
	} else if g.cursor.CanMoveUp() {
		g.cursor.MoveUp()
	} else if g.cursor.CanMoveDown() {
		g.cursor.MoveDown()
	} else if g.cursor.CanMoveBack() {
		g.cursor.MoveBack()

	} else {
		if len(g.possiblePaths) > 0 {
			fmt.Println("Game over")
			// Reset cursor
			g.cursor.Reset()
			// Next possible path
			g.SetNextPossiblePath()
			// Is Game over
		} else {
			g.isOver = true
		}
	}

	if g.cursor.IsCurrentPositionGoal() {
		g.cursor.IncreaseGoal()
	}
}

func (g *Game) GetValue(x, y int) *Value {
	if x < 0 || y < 0 || x >= len(g.grid[0]) || y >= len(g.grid) {
		return nil
	}
	return g.grid[y][x]
}

func (g *Game) GetTotalExplored() int {
	total := 0
	for _, row := range g.grid {
		for _, val := range row {
			total += val.GetExplored()
		}
	}
	return total
}

func (g *Game) GetGrid() [][]*Value {
	return g.grid
}

func (g *Game) GetCursor() *Cursor {
	return g.cursor
}

func (g *Game) GetCheckedPaths() []Position {
	return g.checkedPaths
}
