package main

import (
	"flag"
	"log"
	"strings"

	"github.com/tryy3/advent-of-code/day-4-part-2/gui"
	"github.com/tryy3/advent-of-code/day-4-part-2/process"
)

func main() {
	file := flag.String("file", "input.txt", "input file")
	input := flag.String("input", "", "input")
	flag.Parse()

	var processor *process.Processor
	var err error

	if *input != "" {
		processor, err = process.LoadProcessorFromReader(strings.NewReader(*input))
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		processor, err = process.LoadProcessorFromFile(*file)
		if err != nil {
			log.Println(err)
			return
		}
	}

	gui := gui.NewGui(processor)
	gui.Run()
}
