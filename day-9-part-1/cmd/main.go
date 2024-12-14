package main

import (
	"flag"
	"log"
	"strings"

	"github.com/tryy3/advent-of-code/day-9-part-1/process"
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

	processor.Run()
}
