package main

import (
	"flag"
	"log"
	"strings"

	"github.com/tryy3/advent-of-code/day-7-part-1/calculation"
)

func main() {
	file := flag.String("file", "input.txt", "input file")
	input := flag.String("input", "", "input")
	flag.Parse()

	var calculator *calculation.Calculator
	var err error

	if *input != "" {
		calculator, err = calculation.LoadCalculatorFromReader(strings.NewReader(*input))
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		calculator, err = calculation.LoadCalculatorFromFile(*file)
		if err != nil {
			log.Println(err)
			return
		}
	}

	total := calculator.Calculate()
	log.Println(total)
}
