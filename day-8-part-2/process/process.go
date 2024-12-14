package process

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Processor struct {
	grid [][]string

	symbols map[string][]*Symbol

	uniqueAntiNodes map[string]bool
}

func NewProcessor() *Processor {
	return &Processor{
		grid:            make([][]string, 0),
		symbols:         make(map[string][]*Symbol),
		uniqueAntiNodes: make(map[string]bool),
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
	processor := NewProcessor()
	buffer := bufio.NewReader(reader)

	x := 0
	y := 0

	for {
		r, _, err := buffer.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error reading file: %v", err)
		}
		if r == '\n' {
			x = 0
			y++
			continue
		} else if r == '\r' {
			continue
		}

		processor.PlaceOnGrid(string(r), x, y)

		if r != '.' {
			value := string(r)
			processor.symbols[value] = append(processor.symbols[value], NewSymbol(value, x, y))
		}

		x++
	}

	return processor, nil
}

func (p *Processor) PlaceOnGrid(value string, x, y int) {
	if len(p.grid)-1 < y {
		p.grid = append(p.grid, make([]string, 0))
	}

	if len(p.grid[y])-1 < x {
		p.grid[y] = append(p.grid[y], value)
	} else {
		p.grid[y][x] = value
	}
}

func (p *Processor) PaintCanvas() string {
	builder := strings.Builder{}

	for _, rows := range p.grid {
		for _, cell := range rows {
			builder.WriteString(cell)
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func (p *Processor) IsInMap(x, y int) bool {
	if y < 0 || y >= len(p.grid) {
		return false
	}

	if x < 0 || x >= len(p.grid[y]) {
		return false
	}

	return true
}

func (p *Processor) AddUniqueAntiNode(key string, position string) {

	p.uniqueAntiNodes[position] = true
}

func (p *Processor) PlaceAntiNode(node1, node2 *Symbol) {
	nx := (node1.GetX() * 2) - node2.GetX()
	ny := (node1.GetY() * 2) - node2.GetY()

	if !p.IsInMap(nx, ny) {
		return
	}

	p.AddUniqueAntiNode(node1.String(), fmt.Sprintf("x:%d;y:%d", nx, ny))
	p.PlaceOnGrid("#", nx, ny)
}

func (p *Processor) PlaceAntiNode2(node1, node2 *Symbol) {
	nx := (node2.GetX() - node1.GetX())
	ny := (node2.GetY() - node1.GetY())

	for {
		fmt.Printf("nx: %d, ny: %d\n", nx, ny)

		if !p.IsInMap(nx, ny) {
			break
		}

		p.AddUniqueAntiNode(node1.String(), fmt.Sprintf("x:%d;y:%d", nx, ny))
		p.PlaceOnGrid("#", nx, ny)
		nx += node2.GetX()
		ny += node2.GetY()
	}
}

func (p *Processor) PlaceAllAntiNodes() {
	for _, symbols := range p.symbols {
		for _, symbol1 := range symbols {
			for _, symbol2 := range symbols {
				if symbol1.GetID() == symbol2.GetID() {
					continue
				}

				p.PlaceAntiNode(symbol1, symbol2)
				p.PlaceAntiNode(symbol2, symbol1)

				p.PlaceAntiNode2(symbol1, symbol2)
				p.PlaceAntiNode2(symbol2, symbol1)
			}
		}
	}
}

func (p *Processor) PlaceAllNodes() {
	for _, symbols := range p.symbols {
		for _, symbol := range symbols {
			p.PlaceOnGrid(symbol.String(), symbol.GetX(), symbol.GetY())
		}
	}
}

func (p *Processor) GetUniqueAntiNodes() int {
	total := 0

	for _, antiNodes := range p.uniqueAntiNodes {
		if antiNodes {
			total++
		}
	}

	return total
}

func (p *Processor) Run() {
	p.PlaceAllAntiNodes()
	p.PlaceAllNodes()

	fmt.Println(p.PaintCanvas())
	fmt.Println(p.GetUniqueAntiNodes())
}
