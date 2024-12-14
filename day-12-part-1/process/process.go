package process

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Processor struct {
	rawMap [][]*Symbol

	grids []*Grid
}

func NewProcessor() *Processor {
	return &Processor{
		grids:  make([]*Grid, 0),
		rawMap: make([][]*Symbol, 0),
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
			x = 0
			y++
		} else if r == '\r' {
			continue
		} else {
			if len(processor.rawMap) <= y {
				processor.rawMap = append(processor.rawMap, make([]*Symbol, 0))
			}
			processor.rawMap[y] = append(processor.rawMap[y], NewSymbol(string(r)))
		}
		x++
	}

	return processor, nil
}

func (p *Processor) TraverseGridBuild(grid *Grid, currentPosition *Position) {
	// Check left
	left := NewPosition(currentPosition.x-1, currentPosition.y)
	if left.x >= 0 && left.y >= 0 && left.x < len(p.rawMap[0]) && left.y < len(p.rawMap) {
		newPos := p.rawMap[left.y][left.x]
		if newPos.IsUnchecked() && newPos.Value() == grid.Value() {
			grid.Set(left.x, left.y, newPos.Value())
			newPos.Check()
			p.TraverseGridBuild(grid, left)
		}
	}
	// Check right
	right := NewPosition(currentPosition.x+1, currentPosition.y)
	if right.x >= 0 && right.y >= 0 && right.x < len(p.rawMap[0]) && right.y < len(p.rawMap) {
		newPos := p.rawMap[right.y][right.x]
		if newPos.IsUnchecked() && newPos.Value() == grid.Value() {
			grid.Set(right.x, right.y, newPos.Value())
			newPos.Check()
			p.TraverseGridBuild(grid, right)
		}
	}
	// Check Up
	up := NewPosition(currentPosition.x, currentPosition.y-1)
	if up.x >= 0 && up.y >= 0 && up.x < len(p.rawMap[0]) && up.y < len(p.rawMap) {
		newPos := p.rawMap[up.y][up.x]
		if newPos.IsUnchecked() && newPos.Value() == grid.Value() {
			grid.Set(up.x, up.y, newPos.Value())
			newPos.Check()
			p.TraverseGridBuild(grid, up)
		}
	}
	// Check Down
	down := NewPosition(currentPosition.x, currentPosition.y+1)
	if down.x >= 0 && down.y >= 0 && down.x < len(p.rawMap[0]) && down.y < len(p.rawMap) {
		newPos := p.rawMap[down.y][down.x]
		if newPos.IsUnchecked() && newPos.Value() == grid.Value() {
			grid.Set(down.x, down.y, newPos.Value())
			newPos.Check()
			p.TraverseGridBuild(grid, down)
		}
	}

	grid.Set(currentPosition.x, currentPosition.y, grid.Value())
}

func (p *Processor) BuildGrids() {
	for y, row := range p.rawMap {
		for x, symbol := range row {
			if symbol.IsUnchecked() {
				newGrid := NewGrid(symbol.Value())
				pos := NewPosition(x, y)
				p.TraverseGridBuild(newGrid, pos)

				p.grids = append(p.grids, newGrid)
			}
		}
	}
}

func (p *Processor) CalculateCost() int {
	totalCost := 0
	for _, grid := range p.grids {
		totalCost += grid.Cost()
	}

	return totalCost
}

func (p *Processor) Run() {
	p.BuildGrids()
	fmt.Printf("Total cost: %d\n", p.CalculateCost())
}
