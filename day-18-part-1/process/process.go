package process

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Processor struct {
	corruption   map[string]bool
	processBytes []*Position
	processCount int
	player       *Position
	maxWidth     int
	maxHeight    int
	path         []Position
}

func NewProcessor(maxWidth, maxHeight int) *Processor {
	return &Processor{
		corruption:   make(map[string]bool),
		processBytes: make([]*Position, 0),
		processCount: 0,
		player:       NewPosition(0, 0),
		maxWidth:     maxWidth,
		maxHeight:    maxHeight,
	}
}

func (p *Processor) LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return p.LoadFromReader(file)
}

func (p *Processor) LoadFromReader(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}
		p.corruption[fmt.Sprintf("%d,%d", x, y)] = false
		p.processBytes = append(p.processBytes, NewPosition(x, y))
	}

	return nil
}

func (p *Processor) Update() {
	// corruptedPosition := p.processBytes[p.processCount]
	// if _, ok := p.corruption[fmt.Sprintf("%d,%d", corruptedPosition.GetX(), corruptedPosition.GetY())]; ok {
	// 	p.corruption[fmt.Sprintf("%d,%d", corruptedPosition.GetX(), corruptedPosition.GetY())] = true
	// 	p.processCount++
	// }

	mapGrid := make([][]bool, p.maxHeight)
	for i := range mapGrid {
		mapGrid[i] = make([]bool, p.maxWidth)
	}

	blockSchedule := make(map[int]map[Position]bool)
	for i := 0; i < 1024; i++ {
		processByte := p.processBytes[i]
		mapGrid[processByte.GetY()][processByte.GetX()] = true
		p.corruption[fmt.Sprintf("%d,%d", processByte.GetX(), processByte.GetY())] = true
	}

	start := Position{x: 0, y: 0}
	goal := Position{x: p.maxWidth - 1, y: p.maxHeight - 1}
	p.path = AStar(start, goal, mapGrid, blockSchedule, 0)

	// Part2, find next byte that will cut the path of
	for i := 0; i < len(p.processBytes); i++ {
		processByte := p.processBytes[i]
		for _, pos := range p.path {
			if pos.GetX() == processByte.GetX() && pos.GetY() == processByte.GetY() {
				fmt.Printf("Byte %d will cut the path at %d,%d\n", i, processByte.GetX(), processByte.GetY())
			}
		}
	}

	if p.path != nil {
		fmt.Println("Path found:", len(p.path)-1)
	} else {
		fmt.Println("No path found.")
	}
}

func (p *Processor) Render() {
	for y := 0; y < p.maxHeight; y++ {
		for x := 0; x < p.maxWidth; x++ {
			var path = false
			for _, pos := range p.path {
				if pos.GetX() == x && pos.GetY() == y {
					path = true
					break
				}
			}

			if p.player.GetX() == x && p.player.GetY() == y {
				fmt.Print("P")
			} else if fallen, ok := p.corruption[fmt.Sprintf("%d,%d", x, y)]; ok && fallen {
				fmt.Print("#")
			} else if path {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
