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
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
	}

	return processor, nil
}

func (p *Processor) Run() {
}
