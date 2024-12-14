package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tryy3/advent-of-code/day-6-part-1/game"
)

func main() {
	g, err := game.LoadGameFromFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(game.NewModel(g))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
