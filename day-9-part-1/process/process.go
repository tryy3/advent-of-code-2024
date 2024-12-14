package process

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Disk struct {
	id    int
	value string
	size  int
}

type Processor struct {
	input  []int
	canvas []Disk
}

func NewProcessor() *Processor {
	return &Processor{
		input:  []int{},
		canvas: []Disk{},
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

	for {
		r, _, err := buffer.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		num, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, fmt.Errorf("error converting rune to int: %v", err)
		}

		processor.input = append(processor.input, num)
	}

	return processor, nil
}

func (p *Processor) Repeat(id int, value string, times int) {
	for i := 0; i < times; i++ {
		disk := Disk{
			id:    id,
			value: value,
			size:  times,
		}
		p.canvas = append(p.canvas, disk)
	}
}

func (p *Processor) Layout() {
	state := 0
	id := 0
	for _, num := range p.input {
		if state == 0 {
			p.Repeat(id, fmt.Sprintf("%d", id), num)
			state = 1
			id++
		} else if state == 1 {
			p.Repeat(-1, ".", num)
			state = 0
		}
	}
}

func (p *Processor) SortLayout() {
	length := len(p.canvas)
	canvas := []Disk{}
	left := -1

	for i := length - 1; i >= 0; i-- {
		disk := p.canvas[i]
		if disk.id == -1 {
			continue
		}

		if left > i {
			break
		}

		for {
			left++
			if left > i {
				break
			}

			leftChar := p.canvas[left]
			if leftChar.id == -1 {
				canvas = append(canvas, disk)
				break
			} else {
				canvas = append(canvas, leftChar)
			}
		}
	}

	p.canvas = canvas
	p.Repeat(-1, ".", length-len(canvas))
}

func (p *Processor) FlipValues(pos1, pos2 int) {
	p.canvas[pos1], p.canvas[pos2] = p.canvas[pos2], p.canvas[pos1]
}

func (p *Processor) CheckFreeSpace(pos int, id int) int {
	space := 0

	for {
		if pos >= len(p.canvas) {
			break
		}

		disk := p.canvas[pos]
		if disk.id == id {
			space++
		} else {
			break
		}

		pos++
	}

	return space
}

func (p *Processor) AttemptToSort(disk Disk, pos int) {
	for i := 0; i < len(p.canvas); i++ {
		freeDisk := p.canvas[i]
		freeSpace := p.CheckFreeSpace(i, freeDisk.id)
		// fmt.Printf("freeSpace: %d; freeDiskSize: %d\n", freeSpace, freeDisk.size)

		if freeDisk.id == -1 && freeSpace >= disk.size {
			if i >= pos-disk.size+1 {
				return
			}

			// freeDisk.id = disk.id
			// freeDisk.size -= disk.size

			for v := 0; v < disk.size; v++ {
				p.FlipValues(i+v, pos-v)
			}

			// p.Print()
			// fmt.Println()

			break
		}
	}
}

func (p *Processor) SortFiles() {
	checkedDisks := map[int]bool{}

	for i := len(p.canvas) - 1; i >= 0; i-- {
		disk := p.canvas[i]
		if disk.id == -1 {
			continue
		}

		if checkedDisks[disk.id] {
			continue
		}

		p.AttemptToSort(disk, i)
		checkedDisks[disk.id] = true
	}
}

func (p *Processor) CalculateSum() int {
	sum := 0

	pos := 0
	for i := 0; i < len(p.canvas); i++ {
		disk := p.canvas[i]
		if disk.id == -1 {
			pos++
			continue
		}

		sum += pos * disk.id

		pos++
	}

	return sum
}

func (p Processor) Print() {
	for _, disk := range p.canvas {
		fmt.Printf("%s", disk.value)
	}
}

func (p *Processor) Run() {
	p.Layout()
	// p.Print()
	// fmt.Println()
	p.SortFiles()
	// p.Print()
	// fmt.Println()

	sum := p.CalculateSum()
	fmt.Println(sum)
}
