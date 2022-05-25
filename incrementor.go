package drvolodko3

import (
	"math"
)

type Incrementor interface {
	// GetNumber returns current value.
	GetNumber() int
	// IncrementNumber increases current value by 1 when current value is less than maximum value,
	// sets current value to 0 otherwise.
	IncrementNumber()
	// SetMaximumValue changes maximum value, panics if maximumValue is negative,
	// sets current value to maximum value if current value greater than maximum value.
	SetMaximumValue(maximumValue int)
}

// NewIncrementor returns an Incrementor with maximum value equal to math.MaxInt.
func NewIncrementor() Incrementor {
	return &incrementor{
		maximumValue: math.MaxInt,
	}
}

// GetNumber implements Incrementor.GetNumber.
func (i *incrementor) GetNumber() int {
	return i.currentValue
}

// IncrementNumber implements Incrementor.IncrementNumber.
func (i *incrementor) IncrementNumber() {
	if i.currentValue < i.maximumValue {
		i.currentValue++
	} else {
		i.currentValue = 0
	}
}

// SetMaximumValue implements Incrementor.SetMaximumValue.
func (i *incrementor) SetMaximumValue(maximumValue int) {
	if maximumValue < 0 {
		panic("negative maximum value")
	}
	i.maximumValue = maximumValue
	if i.currentValue > maximumValue {
		i.currentValue = maximumValue
	}
}

type incrementor struct {
	currentValue int
	maximumValue int
}
