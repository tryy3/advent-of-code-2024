package process

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type Processor struct {
	claws []*Claw
}

func NewProcessor() *Processor {
	return &Processor{
		claws: make([]*Claw, 0),
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
	reg := regexp.MustCompile(`(?:Button A: X\+(\d+), Y\+(\d+)\r\nButton B: X\+(\d+), Y\+(\d+)\r\nPrize: X=(\d+), Y=(\d+))`)

	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	matches := reg.FindAllStringSubmatch(string(content), -1)

	for _, match := range matches {
		clawAPosA := parseInt(match[1])
		clawAPosB := parseInt(match[2])
		clawA := NewButton(clawAPosA, clawAPosB, 3)

		clawBPosA := parseInt(match[3])
		clawBPosB := parseInt(match[4])
		clawB := NewButton(clawBPosA, clawBPosB, 1)

		prizePosA := parseInt(match[5]) + 10000000000000
		prizePosB := parseInt(match[6]) + 10000000000000
		prize := NewPosition(prizePosA, prizePosB)

		processor.claws = append(processor.claws, NewClaw(clawA, clawB, prize))
	}

	return processor, nil
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func (p *Processor) Run() {
	for _, claw := range p.claws {
		for {
			claw.MoveClawBForward()
			if claw.IsPositionFurtherThanPrize() {
				break
			}
		}

		if claw.IsPrizeReached() {
			break
		}

		for {
			for {
				claw.MoveClawBBackward()
				if claw.IsPositionBeforePrize() {
					break
				}
			}

			if claw.IsPrizeReached() {
				break
			}

			for {
				claw.MoveClawAForward()
				if claw.IsPositionFurtherThanPrize() {
					break
				}
			}

			if claw.IsPrizeReached() {
				break
			}

			if claw.GetTotalButtonBCount() < 0 {
				break
			}
		}
	}

	total := 0
	for _, claw := range p.claws {
		if claw.IsPrizeReached() {
			fmt.Printf("Button A: %d, Button B: %d, Prize Position: %v\n", claw.GetTotalButtonACount(), claw.GetTotalButtonBCount(), claw.prize)
			total += claw.GetTotalCost()
		}
	}

	fmt.Printf("Total: %d\n", total)
}
