package process

type Program struct {
	instructions []int
}

func NewProgram() *Program {
	return &Program{instructions: []int{}}
}

func (p *Program) AddInstruction(instruction int) {
	p.instructions = append(p.instructions, instruction)
}

func (p *Program) ReadInstructionAt(index int) (opcode, operand int, ok bool) {
	if index >= len(p.instructions) {
		return 0, 0, false
	}

	opcode = p.instructions[index]
	operand = p.instructions[index+1]
	return opcode, operand, true
}
