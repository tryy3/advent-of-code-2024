package gui

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tryy3/advent-of-code/day-10-part-1/game"
)

var (
	fontSize  int32 = 16
	fontWidth int32 = 18
	offsetX   int32 = 10
	offsetY   int32 = 10
)

var (
	DefaultColor    = rl.White
	UnexploredColor = rl.Maroon
	ExploredColor   = rl.Green
	CheckedColor    = rl.White
	RouteColor      = rl.Yellow
	CursorColor     = rl.Blue
)

type Gui struct {
	game *game.Game
}

func NewGui(game *game.Game) *Gui {
	return &Gui{game: game}
}

func (g *Gui) Run() {
	height := 1020
	rl.InitWindow(1215, int32(height), "Advent of Code 2024 - Day 7 Part 2")
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		g.drawGrid()

		rl.DrawText(fmt.Sprintf("Total Distinct Routes Explored: %d", g.game.GetTotalExplored()), offsetX, int32(height)-fontSize, fontSize, rl.White)

		rl.EndDrawing()

		if !g.game.IsOver() {
			g.game.Update()
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func (g *Gui) drawGrid() {
	grid := g.game.GetGrid()
	for y, row := range grid {
		for x, val := range row {
			var rawValue int = val.GetValue()
			var parsedValue string
			var color rl.Color

			if rawValue == -1 {
				parsedValue = "."
			} else {
				parsedValue = fmt.Sprintf("%d", rawValue)
			}

			if rawValue == 0 {
				color = UnexploredColor
			} else {
				color = DefaultColor
			}

			cursor := g.game.GetCursor()

			checkedPaths := g.game.GetCheckedPaths()
			for _, p := range checkedPaths {
				if p.GetX() == x && p.GetY() == y {
					color = CheckedColor
				}
			}

			route := cursor.GetCurrentRoute()
			for _, r := range route {
				if r.GetPosition().GetX() == x && r.GetPosition().GetY() == y {
					color = RouteColor
				}
			}

			cursorPosition := cursor.GetCurrentPosition()
			if cursor != nil && cursorPosition != nil && cursorPosition.GetX() == x && cursorPosition.GetY() == y {
				color = CursorColor
			}

			if val.GetExplored() > 0 {
				color = ExploredColor
			}

			rl.DrawText(parsedValue, (int32(x)*fontWidth)+offsetX, (int32(y)*fontSize)+offsetY, fontSize, color)
		}
	}
}
