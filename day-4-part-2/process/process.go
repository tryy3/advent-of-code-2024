package process

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Processor struct {
	grid [][]*Symbol
}

func NewProcessor() *Processor {
	return &Processor{
		grid: [][]*Symbol{},
	}
}

func LoadProcessorFromFile(path string) (*Processor, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return LoadProcessorFromReader(file)
}

func LoadProcessorFromReader(reader io.Reader) (*Processor, error) {
	processor := &Processor{}
	buffer := bufio.NewReader(reader)

	var x, y int

	for {
		r, _, err := buffer.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		if r == '\n' {
			y++
			x = 0
			continue
		}
		if r == '\r' {
			continue
		}
		symbol := NewSymbol(string(r), NewPosition(x, y))
		if len(processor.grid) < y+1 {
			processor.grid = append(processor.grid, []*Symbol{})
		}
		processor.grid[y] = append(processor.grid[y], symbol)
		x++
	}

	return processor, nil
}

func (p *Processor) Run() {
	for i := 0; i < 1; i++ {
		p.Update()
	}
}

func (p *Processor) GetSymbolValue(x, y int) string {
	if x < 0 || y < 0 || y >= len(p.grid) || x >= len(p.grid[y]) {
		return ""
	}
	return p.grid[y][x].GetSymbol()
}

func (p *Processor) FindOccurences(symbol *Symbol) int {
	if symbol.GetSymbol() != "A" {
		return 0
	}

	var word string
	occurence := 0
	x, y := symbol.GetPosition().GetX(), symbol.GetPosition().GetY()

	// Check Left top to bottom right
	word = fmt.Sprintf("%s%s%s", p.GetSymbolValue(x-1, y-1), symbol.GetSymbol(), p.GetSymbolValue(x+1, y+1))
	if word == "MAS" {
		occurence++
	}
	// check Left bottom to top right
	word = fmt.Sprintf("%s%s%s", p.GetSymbolValue(x-1, y+1), symbol.GetSymbol(), p.GetSymbolValue(x+1, y-1))
	if word == "MAS" {
		occurence++
	}
	// Check Right top to bottom left
	word = fmt.Sprintf("%s%s%s", p.GetSymbolValue(x+1, y-1), symbol.GetSymbol(), p.GetSymbolValue(x-1, y+1))
	if word == "MAS" {
		occurence++
	}
	// Check Right bottom to top left
	word = fmt.Sprintf("%s%s%s", p.GetSymbolValue(x+1, y+1), symbol.GetSymbol(), p.GetSymbolValue(x-1, y-1))
	if word == "MAS" {
		occurence++
	}
	return occurence
}

func (p *Processor) Update() {
	total := 0
	for y := 1; y < len(p.grid)-1; y++ {
		for x := 1; x < len(p.grid[y])-1; x++ {
			symbol := p.grid[y][x]
			if symbol.GetSymbol() == "." {
				continue
			}
			occurence := p.FindOccurences(symbol)
			if occurence >= 2 {
				// total += occurence
				total++
			}
			if occurence > 0 {
				fmt.Printf("%s", symbol.GetSymbol())
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println(total)
}

func (p *Processor) GetGrid() [][]*Symbol {
	return p.grid
}
