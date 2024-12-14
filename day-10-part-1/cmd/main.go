package main

import (
	"flag"
	"log"
	"strings"

	"github.com/tryy3/advent-of-code/day-10-part-1/game"
	"github.com/tryy3/advent-of-code/day-10-part-1/gui"
)

func main() {
	file := flag.String("file", "file.txt", "input file")
	input := flag.String("input", "", "input")
	flag.Parse()

	var g *game.Game
	var err error

	if *input != "" {
		g, err = game.LoadGameFromReader(strings.NewReader(*input))
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		g, err = game.LoadGameFromFile(*file)
		if err != nil {
			log.Println(err)
			return
		}
	}

	gui := gui.NewGui(g)
	gui.Run()

	// for {
	// 	if g.IsOver() {
	// 		break
	// 	}
	// 	g.Update()
	// 	// time.Sleep(100 * time.Millisecond)
	// }

	// fmt.Println(g.GetTotalExplored())
}
