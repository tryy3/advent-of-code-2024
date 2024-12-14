package calculation

type Value struct {
	OriginalValue   int
	CalculatedValue int
	Calculated      bool
}

func NewValue(value int) *Value {
	return &Value{
		OriginalValue:   value,
		CalculatedValue: 0,
		Calculated:      false,
	}
}

func (v *Value) GetOriginalValue() int {
	return v.OriginalValue
}

func (v *Value) GetValue() int {
	if v.Calculated {
		return v.CalculatedValue
	}
	return v.OriginalValue
}

func (v *Value) SetCalculatedValue(value int) {
	v.CalculatedValue = value
	v.Calculated = true
}

func (v *Value) Reset() {
	v.CalculatedValue = 0
	v.Calculated = false
}

func (v *Value) Calculate(firstValue *Value, operator *Operator) {
	v.SetCalculatedValue(operator.Calculate(firstValue.GetValue(), v.GetValue()))
}
