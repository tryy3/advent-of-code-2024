package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/tryy3/advent-of-code/day-11-part-1-map/process"
	_ "modernc.org/sqlite"
)

func main() {
	// Create memory profile file
	f, err := os.Create("mem.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	file := flag.String("file", "input.txt", "input file")
	input := flag.String("input", "", "input")
	generations := flag.Int("generations", 75, "number of generations to blink")
	flag.Parse()

	var processor *process.Processor

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

	currentTime := time.Now()
	fmt.Println(processor.GenerateValues(*generations))
	elapsedTime := time.Since(currentTime)
	fmt.Printf("Blinking generation %d took %s\n", *generations, elapsedTime)

	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal(err)
	}
}
