package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/tryy3/advent-of-code/day-15-part-1/process"
)

func main() {
	file := flag.String("file", "test.txt", "input file")
	input := flag.String("input", "", "input")
	steps := flag.Int("steps", 15, "number of steps")
	flag.Parse()

	var err error
	processor := process.NewProcessor()

	if *input != "" {
		err = processor.LoadFromReader(strings.NewReader(*input))
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		err = processor.LoadFromFile(*file)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if *steps <= 0 {
		*steps = processor.GetStepCount()
	}

	for i := 0; i < *steps; i++ {
		processor.Update()
	}

	fmt.Println(processor.GetSum())
}
