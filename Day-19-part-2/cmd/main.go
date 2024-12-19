package main

import (
	"flag"
	"log"
	"strings"

	"github.com/tryy3/advent-of-code/day-8-part-1/process"
)

func main() {
	file := flag.String("file", "input.txt", "input file")
	input := flag.String("input", "", "input")
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

	processor.Update()
}
