package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

type Position struct {
	x      int
	y      int
	letter string
}

type Grid struct {
	grid   [][]string
	result []Position
}

func main() {
	file, err := openFile()
	if err != nil {
		log.Fatalf("error executing openFile: %v", err)
	}
	defer file.Close()

	grid := NewGrid()
	err = grid.ReadData(file)
	if err != nil {
		log.Fatalf("error executing ReadData: %v", err)
	}

	horizontalCount := grid.CountHorizontal()
	log.Printf("Horizontal count: %d\n", horizontalCount)

	verticalCount := grid.CountVertical()
	log.Printf("Vertical count: %d\n", verticalCount)

	backwardsCount := grid.CountBackwardsHorizontal()
	log.Printf("Backwards count: %d\n", backwardsCount)

	backwardsVerticalCount := grid.CountBackwardsVertical()
	log.Printf("Backwards vertical count: %d\n", backwardsVerticalCount)

	diagonalTLBRCount := grid.CountDiagonalTLBR()
	log.Printf("Diagonal TLBR count: %d\n", diagonalTLBRCount)

	diagonalTRBLCount := grid.CountDiagonalTRBL()
	log.Printf("Diagonal TRBL count: %d\n", diagonalTRBLCount)

	total := horizontalCount + verticalCount + backwardsCount + backwardsVerticalCount + diagonalTLBRCount + diagonalTRBLCount
	log.Printf("Total count: %d\n", total)

	grid.PrintResult()
}

func NewGrid() *Grid {
	return &Grid{
		grid: [][]string{},
	}
}

func openFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	return file, nil
}

func (g *Grid) ReadData(r io.Reader) error {
	reader := bufio.NewReader(r)

	x := 0
	y := 0

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			// Break at EOF
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("error reading file: %v", err)
		}

		if r == '\n' {
			x = 0
			y++

		} else if unicode.IsLetter(r) {
			if len(g.grid) <= y {
				g.grid = append(g.grid, []string{})
			}
			g.grid[y] = append(g.grid[y], string(r))
			x++
		}
	}

	return nil
}

func (g *Grid) Width() int {
	return len(g.grid[0])
}

func (g *Grid) Height() int {
	return len(g.grid)
}

func (g *Grid) IsResult(x, y int) bool {
	for _, position := range g.result {
		if position.x == x && position.y == y {
			return true
		}
	}
	return false
}

func (g *Grid) PrintResult() {
	height := g.Height()
	width := g.Width()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if g.IsResult(x, y) {
				fmt.Print(g.grid[y][x])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (g *Grid) CountHorizontal() int {
	height := g.Height()
	width := g.Width()
	count := 0

	// Iterate through each column
	for y := 0; y < height; y++ {
		currentWord := ""
		// Build word by going down each row
		for x := 0; x < width; x++ {
			currentWord += g.grid[y][x]

			if !strings.HasPrefix("XMAS", currentWord) {
				currentWord = g.grid[y][x]
			}

			if currentWord == "XMAS" {
				for i := 0; i < 4; i++ {
					g.result = append(g.result, Position{x - 3 + i, y, string(currentWord[i])})
				}

				count++
				currentWord = ""
			}
		}
	}

	return count
}

func (g *Grid) CountVertical() int {
	height := g.Height()
	width := g.Width()
	count := 0

	// Iterate through each column
	for x := 0; x < width; x++ {
		currentWord := ""
		// Build word by going down each row
		for y := 0; y < height; y++ {
			currentWord += g.grid[y][x]

			if !strings.HasPrefix("XMAS", currentWord) {
				currentWord = g.grid[y][x]
			}

			if currentWord == "XMAS" {
				for i := 0; i < 4; i++ {
					g.result = append(g.result, Position{x, y - 3 + i, string(currentWord[i])})
				}

				count++
				currentWord = ""
			}
		}
	}

	return count
}

func (g *Grid) CountBackwardsHorizontal() int {
	height := g.Height()
	width := g.Width()
	count := 0

	// Build word by going down each row
	for y := 0; y < height; y++ {
		currentWord := ""

		for x := width - 1; x >= 0; x-- {
			currentWord += g.grid[y][x]

			if !strings.HasPrefix("XMAS", currentWord) {
				currentWord = g.grid[y][x]
			}

			if currentWord == "XMAS" {
				for i := 0; i < 4; i++ {
					g.result = append(g.result, Position{x + i, y, string(currentWord[i])})
				}

				count++
				currentWord = ""
			}
		}
	}

	return count
}

func (g *Grid) CountBackwardsVertical() int {
	count := 0

	// Get the dimensions of the grid
	if len(g.grid) == 0 {
		return 0
	}
	height := g.Height()
	width := g.Width()

	// Iterate through each column
	for x := 0; x < width; x++ {
		currentWord := ""
		// Build word by going down each row
		for y := height - 1; y >= 0; y-- {
			currentWord += g.grid[y][x]

			if !strings.HasPrefix("XMAS", currentWord) {
				currentWord = g.grid[y][x]
			}

			if currentWord == "XMAS" {
				for i := 0; i < 4; i++ {
					g.result = append(g.result, Position{x, y + i, string(currentWord[i])})
				}

				count++
				currentWord = ""
			}
		}
	}

	return count
}

func (g *Grid) CountDiagonalTLBR() int {
	height := g.Height()
	width := g.Width()
	count := 0

	for y := 0; y < height-3; y++ {
		for x := 0; x < width-3; x++ {
			word := g.grid[y][x] + g.grid[y+1][x+1] + g.grid[y+2][x+2] + g.grid[y+3][x+3]
			if word == "XMAS" {
				g.result = append(g.result, Position{x, y, string(word[0])})
				g.result = append(g.result, Position{x + 1, y + 1, string(word[1])})
				g.result = append(g.result, Position{x + 2, y + 2, string(word[2])})
				g.result = append(g.result, Position{x + 3, y + 3, string(word[3])})

				count++
			}
		}
	}

	for y := 0; y < height-3; y++ {
		for x := width - 1; x >= 3; x-- {
			word := g.grid[y][x] + g.grid[y+1][x-1] + g.grid[y+2][x-2] + g.grid[y+3][x-3]
			if word == "XMAS" {
				g.result = append(g.result, Position{x, y, string(word[0])})
				g.result = append(g.result, Position{x - 1, y + 1, string(word[1])})
				g.result = append(g.result, Position{x - 2, y + 2, string(word[2])})
				g.result = append(g.result, Position{x - 3, y + 3, string(word[3])})

				count++
			}
		}
	}

	return count
}

func (g *Grid) CountDiagonalTRBL() int {
	height := g.Height()
	width := g.Width()
	count := 0

	for y := height - 1; y >= 3; y-- {
		for x := 0; x < width-3; x++ {
			word := g.grid[y][x] + g.grid[y-1][x+1] + g.grid[y-2][x+2] + g.grid[y-3][x+3]
			if word == "XMAS" {
				g.result = append(g.result, Position{x, y, string(word[0])})
				g.result = append(g.result, Position{x + 1, y - 1, string(word[1])})
				g.result = append(g.result, Position{x + 2, y - 2, string(word[2])})
				g.result = append(g.result, Position{x + 3, y - 3, string(word[3])})

				count++
			}
		}
	}

	for y := height - 1; y >= 3; y-- {
		for x := width - 1; x >= 3; x-- {
			word := g.grid[y][x] + g.grid[y-1][x-1] + g.grid[y-2][x-2] + g.grid[y-3][x-3]
			if word == "XMAS" {
				g.result = append(g.result, Position{x, y, string(word[0])})
				g.result = append(g.result, Position{x - 1, y + 1, string(word[1])})
				g.result = append(g.result, Position{x - 2, y + 2, string(word[2])})
				g.result = append(g.result, Position{x - 3, y + 3, string(word[3])})

				count++
			}
		}
	}

	return count
}
