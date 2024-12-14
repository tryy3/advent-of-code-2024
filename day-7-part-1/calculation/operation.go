package calculation

type Operations int

const (
	Addition Operations = iota
	Multiplication
)

type Operator struct {
	Operation Operations
}

func NewOperator(operation Operations) *Operator {
	return &Operator{
		Operation: operation,
	}
}

func (o *Operator) Calculate(firstValue int, secondValue int) int {
	switch o.Operation {
	case Addition:
		return firstValue + secondValue
	case Multiplication:
		return firstValue * secondValue
	}
	return 0
}

func (o *Operator) SwitchOperation() {
	o.Operation = (o.Operation + 1)
	if o.Operation > Multiplication {
		o.Operation = Addition
	}
}

func (o *Operator) Reset() {
	o.Operation = Addition
}

func (o *Operator) String() string {
	switch o.Operation {
	case Addition:
		return "+"
	case Multiplication:
		return "*"
	}
	return ""
}
