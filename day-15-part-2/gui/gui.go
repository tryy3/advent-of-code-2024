package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Gui struct {
}

func NewGui() *Gui {
	return &Gui{}
}

func (g *Gui) Run() {
	rl.InitWindow(1280, 720, "Advent of Code 2024 - Day 7 Part 2")
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}
