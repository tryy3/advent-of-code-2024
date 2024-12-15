package process

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Processor struct {
}

func NewProcessor() *Processor {
	return &Processor{}
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
		log.Println(line)
	}

	return nil
}

func (p *Processor) Update() {
}
