package main

import (
	"flag"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tryy3/advent-of-code/day-6-part-2/game"
)

func main() {
	// Flag for file input
	inputFile := flag.String("input", "../file.txt", "input file")

	// flag for default skip updates
	skipUpdates := flag.Int("skip", 50, "skip updates")

	// flag for tick rate
	tickRate := flag.Duration("tick", time.Millisecond*50, "tick rate")
	flag.Parse()

	g, err := game.LoadGameFromFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	m := game.NewModel(g)
	m.SetSkipCount(*skipUpdates)
	m.SetTickRate(*tickRate)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
