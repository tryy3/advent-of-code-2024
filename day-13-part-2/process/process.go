package process

import (
	"fmt"
	"io"
	"math"
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
		// prizePosA := parseInt(match[5])
		// prizePosB := parseInt(match[6])
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

func calculateDeterminant(buttonA, buttonB *Button) float64 {
	return (float64(buttonA.addX) * float64(buttonB.addY)) - (float64(buttonA.addY) * float64(buttonB.addX))
}

func calculateButtonACount(buttonB *Button, prize *Position, determinant float64) float64 {
	return (float64(buttonB.addY)*float64(prize.x) - float64(buttonB.addX)*float64(prize.y)) / determinant
}

func calculateButtonBCount(buttonA *Button, prize *Position, determinant float64) float64 {
	return (float64(buttonA.addX)*float64(prize.y) - float64(buttonA.addY)*float64(prize.x)) / determinant
}

func (p *Processor) Run() {
	fmt.Println("Start")
	for _, claw := range p.claws {
		// aX * bY - aY * bX
		determinant := calculateDeterminant(claw.buttonA, claw.buttonB)
		fmt.Printf("Determinant: %d\n", int(determinant))

		// (bY * tX - bX * tY) / determinant
		buttonACount := calculateButtonACount(claw.buttonB, claw.prize, determinant)

		// (aX * tY - aY * tX) / determinant
		buttonBCount := calculateButtonBCount(claw.buttonA, claw.prize, determinant)

		fmt.Printf("Button A Count: %d, Button B Count: %d\n", int(buttonACount), int(buttonBCount))

		if buttonACount < 0 || buttonBCount < 0 {
			continue
		}

		fmt.Printf("mod button a: %d\n", math.Mod(buttonACount, 1))
		fmt.Printf("mod button b: %d\n", math.Mod(buttonBCount, 1))

		if math.Mod(buttonACount, 1) != 0 || math.Mod(buttonBCount, 1) != 0 {
			continue
		}

		claw.MoveClawAForward(int(buttonACount))
		claw.MoveClawBForward(int(buttonBCount))
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
