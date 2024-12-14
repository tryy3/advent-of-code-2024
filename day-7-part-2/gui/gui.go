package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tryy3/advent-of-code/day-7-part-2/calculation"
)

type Gui struct {
	calculator *calculation.Calculator
}

func NewGui(calculator *calculation.Calculator) *Gui {
	return &Gui{calculator: calculator}
}

func (g *Gui) Run() {
	rl.InitWindow(1280, 720, "Advent of Code 2024 - Day 7 Part 2")
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		var marginY int32 = 5
		var marginX int32 = 5
		var fontSize int32 = 18
		var rowHeight int32 = 16
		var maxRows int32 = 85
		var rowWidth int32 = 270

		var y int32 = 0
		var x int32 = 0

		for _, row := range g.calculator.Calculation {
			if y >= maxRows {
				y = 0
				x++
			}

			posX := marginX + (x * rowWidth)
			posY := marginY + (y * rowHeight)

			rl.DrawText(row.String(), posX, posY, fontSize, rl.White)

			y++
		}
		rl.EndDrawing()
	}
}
