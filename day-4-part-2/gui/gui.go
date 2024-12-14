package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tryy3/advent-of-code/day-4-part-2/process"
)

var (
	fontSize = 10
	height   = fontSize * 140
	width    = fontSize * 140
)

type Gui struct {
	processor *process.Processor
}

func NewGui(processor *process.Processor) *Gui {
	return &Gui{processor}
}

func (g *Gui) Run() {
	rl.InitWindow(int32(width), int32(height), "Advent of Code 2024 - Day 4 Part 2")
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		g.drawGrid()

		rl.EndDrawing()
	}
}

func (g *Gui) drawGrid() {
	grid := g.processor.GetGrid()
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			symbol := grid[y][x]
			rl.DrawText(symbol.GetSymbol(), int32(x*fontSize), int32(y*fontSize), int32(fontSize), rl.White)
		}
	}
}
