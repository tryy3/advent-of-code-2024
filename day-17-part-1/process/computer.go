package process

import (
	"log"
	"math"
)

type Computer struct {
	registerA *Register
	registerB *Register
	registerC *Register

	program *Program

	out                []int
	instructionPointer int
}

func NewComputer() *Computer {
	return &Computer{
		registerA:          NewRegister(0),
		registerB:          NewRegister(0),
		registerC:          NewRegister(0),
		program:            NewProgram(),
		instructionPointer: 0,
		out:                []int{},
	}
}

func (c *Computer) GetRegisterAValue() int {
	return c.registerA.GetValue()
}

func (c *Computer) GetRegisterBValue() int {
	return c.registerB.GetValue()
}

func (c *Computer) GetRegisterCValue() int {
	return c.registerC.GetValue()
}

func (c *Computer) ReadNextInstruction() bool {
	opcode, operand, ok := c.program.ReadInstructionAt(c.instructionPointer)
	if !ok {
		return false
	}

	jumpInstructions := true

	switch opcode {
	case 0:
		c.run_adv(operand)
	case 1:
		c.run_bxl(operand)
	case 2:
		c.run_bst(operand)
	case 3:
		jumpInstructions = c.run_jnz(operand)
	case 4:
		c.run_bxc(operand)
	case 5:
		c.run_out(operand)
	case 6:
		c.run_bdv(operand)
	case 7:
		c.run_cdv(operand)
	}

	if jumpInstructions {
		c.instructionPointer += 2
	}

	return true
}

func (c *Computer) GetInstructionPointer() int {
	return c.instructionPointer
}

func (c *Computer) GetOutput() []int {
	return c.out
}

func (c *Computer) getCombo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.registerA.GetValue()
	case 5:
		return c.registerB.GetValue()
	case 6:
		return c.registerC.GetValue()
	default:
		log.Fatalf("invalid combo operand: %d", operand)
		return -1
	}
}

// Opcode 0, division
// numerator = Value of register A
// denominator = 2** combo operand
// truncate value then save to A
func (c *Computer) run_adv(operand int) {
	numerator := c.registerA.GetValue()

	combo := c.getCombo(operand)
	denominator := int(math.Pow(2, float64(combo)))

	result := numerator / denominator
	c.registerA.SetValue(result)
}

// Opcode 1, XOR of B and literal operand
func (c *Computer) run_bxl(operand int) {
	literal := operand
	bValue := c.registerB.GetValue()

	result := bValue ^ literal
	c.registerB.SetValue(result)
}

// Opcode 2, Combo operand % 8
// Save to B
func (c *Computer) run_bst(operand int) {
	combo := c.getCombo(operand)
	combo %= 8
	c.registerB.SetValue(combo)
}

// Opcode 3, Jump if not zero
// Jump to literal operand
// don't increase instruction pointer
func (c *Computer) run_jnz(operand int) bool {
	valueA := c.registerA.GetValue()
	if valueA == 0 {
		return true
	}

	c.instructionPointer = operand
	return false
}

// Opcode 4, XOR of B and C
// Save to B
// Read operand but skip value
func (c *Computer) run_bxc(operand int) {
	valueB := c.registerB.GetValue()
	valueC := c.registerC.GetValue()

	result := valueB ^ valueC
	c.registerB.SetValue(result)
}

// Opcode 5, Combo operand % 8
// save to out
func (c *Computer) run_out(operand int) {
	combo := c.getCombo(operand) % 8
	c.out = append(c.out, combo)
}

// Opcode 6, division - adv
// save value to B
func (c *Computer) run_bdv(operand int) {
	numerator := c.registerA.GetValue()

	combo := c.getCombo(operand)
	denominator := int(math.Pow(2, float64(combo)))

	result := numerator / denominator
	c.registerB.SetValue(result)
}

// Opcode 7, division - adv
// save value to C
func (c *Computer) run_cdv(operand int) {
	numerator := c.registerA.GetValue()

	combo := c.getCombo(operand)
	denominator := int(math.Pow(2, float64(combo)))

	result := numerator / denominator
	c.registerC.SetValue(result)
}
