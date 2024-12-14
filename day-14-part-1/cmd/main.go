package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/tryy3/advent-of-code/day-14-part-1/process"
)

func main() {
	file := flag.String("file", "input.txt", "input file")
	maxWidth := flag.Int("max-width", 11, "max width")
	maxHeight := flag.Int("max-height", 7, "max height")
	input := flag.String("input", "", "input")
	saveFolder := flag.String("save-folder", "cmd/images", "save folder")
	loops := flag.Int("loops", 10, "loops")
	flag.Parse()

	var err error

	processor := process.NewProcessor(*maxWidth, *maxHeight, *saveFolder)

	if *input != "" {
		err = processor.LoadReader(strings.NewReader(*input))
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		err = processor.LoadFile(*file)
		if err != nil {
			log.Println(err)
			return
		}
	}

	processor.Clear()
	processor.Render()
	for i := 0; i < *loops; i++ {
		processor.Update()
		processor.Render()
	}
	processor.Render()

	total := processor.CalculateQuadrants()
	fmt.Println(total)
}
