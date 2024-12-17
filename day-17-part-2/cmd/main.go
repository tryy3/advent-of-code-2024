package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tryy3/advent-of-code/day-8-part-1/process"
)

func main() {
	// file := flag.String("file", "file.txt", "input file")
	// input := flag.String("input", "", "input")
	// flag.Parse()

	// var err error
	// processor := process.NewProcessor()

	// if *input != "" {
	// 	err = processor.LoadFromReader(strings.NewReader(*input))
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// } else {
	// 	err = processor.LoadFromFile(*file)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// }

	// for {
	// 	ok := processor.ReadInstruction()
	// 	if !ok {
	// 		break
	// 	}
	// }
	// spew.Dump(processor)
	// processor.PrintOutput()

	// var output = `0,3,5,4,3,0` // Test
	// var output = `7,1,5,2,4,0,7,6,1` // Test

	for i := 0; i < 1<<10; i++ {
		// for i := 2024; i < 2830; i++ {
		processor := process.NewProcessor()
		// 		input := fmt.Sprintf(`Register A: %d
		// Register B: 0
		// Register C: 0

		// Program: 0,3,5,4,3,0`, i) // Test
		input := fmt.Sprintf(`Register A: %d
Register B: 0
Register C: 0

Program: 2,4,1,2,7,5,1,3,4,4,5,5,0,3,3,0`, i) // Test
		// fmt.Printf("input: %s\n", input)
		processor.LoadFromReader(strings.NewReader(input))
		// spew.Dump(processor)
		for {
			ok := processor.ReadInstruction()
			if !ok {
				break
			}
		}

		newOutput := processor.GetOutput()
		outputString := processor.GetOutputString()
		if 7 == newOutput[0] {
			// possibilites.add(aCandidate);
			binaryString := strconv.FormatInt(int64(i), 2)
			binaryString = strings.Repeat("0", 10-len(binaryString)) + binaryString
			fmt.Printf("candidate: %d, binary: %s, result: %s\n", i, binaryString, outputString)
		}
	}
}
