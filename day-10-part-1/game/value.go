package game

import "fmt"

type Value struct {
	position *Position
	value    int
	explored int
}

func NewValue(position *Position, value int) *Value {
	return &Value{position, value, 0}
}

func (v *Value) GetPosition() *Position {
	return v.position
}

func (v *Value) GetValue() int {
	return v.value
}

func (v *Value) String() string {
	return fmt.Sprintf("%d", v.value)
}

func (v *Value) GetExplored() int {
	return v.explored
}

func (v *Value) IncreaseExplored() {
	v.explored++
}
