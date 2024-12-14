package process

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"regexp"
	"strconv"
)

type Processor struct {
	maxWidth   int
	maxHeight  int
	guards     []*Guard
	updates    int
	saveFolder string
}

func NewProcessor(maxWidth, maxHeight int, saveFolder string) *Processor {
	return &Processor{maxWidth, maxHeight, []*Guard{}, 0, saveFolder}
}
func (p *Processor) LoadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return p.LoadReader(file)
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func (p *Processor) LoadReader(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	reg := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := reg.FindStringSubmatch(line)

		position := NewPosition(
			parseInt(matches[1]),
			parseInt(matches[2]),
		)
		velocity := NewVelocity(
			parseInt(matches[3]),
			parseInt(matches[4]),
		)
		guard := NewGuard(position, velocity)
		p.guards = append(p.guards, guard)
	}

	return nil
}

func (p *Processor) Update() {
	p.updates++
	for _, guard := range p.guards {
		x, y := guard.GetPosition().GetX(), guard.GetPosition().GetY()

		newX := (x + guard.GetVelocity().GetX() + p.maxWidth) % p.maxWidth
		newY := (y + guard.GetVelocity().GetY() + p.maxHeight) % p.maxHeight

		guard.GetPosition().SetX(newX)
		guard.GetPosition().SetY(newY)
	}
}

func (p *Processor) GetGuardsAt(x, y int) []*Guard {
	guards := []*Guard{}
	for _, guard := range p.guards {
		if guard.GetPosition().GetX() == x && guard.GetPosition().GetY() == y {
			guards = append(guards, guard)
		}
	}
	return guards
}

func (p *Processor) Render() {
	cubeSize := 10

	// Create a new blank image with a white background
	width, height := p.maxWidth*cubeSize, p.maxHeight*cubeSize
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	fmt.Println("Render", p.updates)
	for y := 0; y < p.maxHeight; y++ {
		for x := 0; x < p.maxWidth; x++ {
			guards := p.GetGuardsAt(x, y)
			if len(guards) > 0 {
				// fmt.Print(len(guards))
				draw.Draw(img, image.Rect(x*cubeSize, y*cubeSize, (x+1)*cubeSize, (y+1)*cubeSize), &image.Uniform{black}, image.Point{}, draw.Src)
			} else {
				// fmt.Print(".")
			}
		}
		// fmt.Println()
	}
	// fmt.Println()

	file, err := os.Create(fmt.Sprintf("%s/render-%d.png", p.saveFolder, p.updates))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func (p *Processor) GetGuardsInQuadrant(fromX, fromY, toX, toY int) int {
	total := 0
	for _, guard := range p.guards {
		if guard.GetPosition().GetX() >= fromX && guard.GetPosition().GetX() <= toX && guard.GetPosition().GetY() >= fromY && guard.GetPosition().GetY() <= toY {
			total++
		}
	}
	return total
}

func (p *Processor) CalculateQuadrants() int {
	topLeftQuadrantFromX := 0
	topLeftQuadrantToX := (p.maxWidth / 2) - 1
	topLeftQuadrantFromY := 0
	topLeftQuadrantToY := (p.maxHeight / 2) - 1
	topLeftQuadrant := p.GetGuardsInQuadrant(topLeftQuadrantFromX, topLeftQuadrantFromY, topLeftQuadrantToX, topLeftQuadrantToY)

	topRightQuadrantFromX := (p.maxWidth / 2) + 1
	topRightQuadrantToX := p.maxWidth - 1
	topRightQuadrantFromY := 0
	topRightQuadrantToY := (p.maxHeight / 2) - 1
	topRightQuadrant := p.GetGuardsInQuadrant(topRightQuadrantFromX, topRightQuadrantFromY, topRightQuadrantToX, topRightQuadrantToY)

	bottomLeftQuadrantFromX := 0
	bottomLeftQuadrantToX := (p.maxWidth / 2) - 1
	bottomLeftQuadrantFromY := (p.maxHeight / 2) + 1
	bottomLeftQuadrantToY := p.maxHeight - 1
	bottomLeftQuadrant := p.GetGuardsInQuadrant(bottomLeftQuadrantFromX, bottomLeftQuadrantFromY, bottomLeftQuadrantToX, bottomLeftQuadrantToY)

	bottomRightQuadrantFromX := (p.maxWidth / 2) + 1
	bottomRightQuadrantToX := p.maxWidth - 1
	bottomRightQuadrantFromY := (p.maxHeight / 2) + 1
	bottomRightQuadrantToY := p.maxHeight - 1
	bottomRightQuadrant := p.GetGuardsInQuadrant(bottomRightQuadrantFromX, bottomRightQuadrantFromY, bottomRightQuadrantToX, bottomRightQuadrantToY)

	fmt.Printf("topLeft from: %d, %d, to: %d, %d, count: %d\n", topLeftQuadrantFromX, topLeftQuadrantFromY, topLeftQuadrantToX, topLeftQuadrantToY, topLeftQuadrant)
	fmt.Printf("topRight from: %d, %d, to: %d, %d, count: %d\n", topRightQuadrantFromX, topRightQuadrantFromY, topRightQuadrantToX, topRightQuadrantToY, topRightQuadrant)
	fmt.Printf("bottomLeft from: %d, %d, to: %d, %d, count: %d\n", bottomLeftQuadrantFromX, bottomLeftQuadrantFromY, bottomLeftQuadrantToX, bottomLeftQuadrantToY, bottomLeftQuadrant)
	fmt.Printf("bottomRight from: %d, %d, to: %d, %d, count: %d\n", bottomRightQuadrantFromX, bottomRightQuadrantFromY, bottomRightQuadrantToX, bottomRightQuadrantToY, bottomRightQuadrant)

	return topLeftQuadrant * topRightQuadrant * bottomLeftQuadrant * bottomRightQuadrant
}

func (p *Processor) Clear() {
	os.RemoveAll(p.saveFolder)
	os.MkdirAll(p.saveFolder, 0755)
}
