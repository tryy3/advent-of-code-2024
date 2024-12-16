package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/tryy3/advent-of-code/day-8-part-1/process"
)

func main() {
	file := flag.String("file", "test.txt", "input file")
	input := flag.String("input", "", "input")
	lowestCost := flag.Int("lowest-cost", 7036, "lowest cost")
	maxSteppedOn := flag.Int("max-stepped-on", 5000, "max stepped on")
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

	processor.SetLowestCost(*lowestCost)
	processor.SetMaxSteppedOn(*maxSteppedOn)
	processor.Begin()
	// processor.Render()
	i := 0
	for !processor.IsOver() {
		processor.Update()
		// processor.Render()
		i++
	}
	// for i := 0; i < 500; i++ {
	// 	processor.Update()
	// 	// processor.Render()
	// }
	fmt.Println(i)
	// processor.Render()
	processor.SumPaths()
}
