package drvolodko3

import (
	"math"
	"testing"
)

// TestIncrementor_GetNumber checks Incrementor.GetNumber
// returns value expected after given number of increments from 0.
func TestIncrementor_GetNumber(t *testing.T) {
	var tests = []struct {
		increments int
		want       int
	}{
		{increments: 0, want: 0},
		{increments: 1, want: 1},
		{increments: 3, want: 3},
	}

	for _, tc := range tests {
		var inc = NewIncrementor()
		for i := 0; i < tc.increments; i++ {
			inc.IncrementNumber()
		}
		var got = inc.GetNumber()
		if got != tc.want {
			t.Errorf("want %d, got %d", tc.want, got)
		}
	}
}

// TestIncrementor_IncrementNumber checks Incrementor.IncrementNumber
// increases current value by a given number of increments from 0
// and sets current value to 0 after maximum value reached.
func TestIncrementor_IncrementNumber(t *testing.T) {
	const maximumNumber = 5
	var tests = []struct {
		increments int
		want       int
	}{
		{increments: 4, want: 4},
		{increments: 5, want: 5},
		{increments: 6, want: 0},
		{increments: 7, want: 1},
		{increments: 8, want: 2},
	}

	for _, tc := range tests {
		var inc = NewIncrementor()
		inc.SetMaximumValue(maximumNumber)
		for i := 0; i < tc.increments; i++ {
			inc.IncrementNumber()
		}

		var got = inc.GetNumber()
		if got != tc.want {
			t.Errorf("want %d, got %d", tc.want, got)
		}
	}
}

// TestIncrementor_SetMaximumValueValid checks Incrementor.SetMaximumValue
// accepts non-negative values.
func TestIncrementor_SetMaximumValueValid(t *testing.T) {
	var tests = []struct {
		maximumNumber int
	}{
		{maximumNumber: 0},
		{maximumNumber: 1},
		{maximumNumber: 5},
		{maximumNumber: math.MaxInt - 1},
		{maximumNumber: math.MaxInt},
	}

	for _, tc := range tests {
		var inc = NewIncrementor()
		inc.SetMaximumValue(tc.maximumNumber)
	}
}

// TestIncrementor_SetMaximumValueInvalid checks Incrementor.SetMaximumValue
// does not accept negative values and panics.
func TestIncrementor_SetMaximumValueInvalid(t *testing.T) {
	var tests = []struct {
		maximumNumber int
	}{
		{maximumNumber: -1},
		{maximumNumber: math.MinInt + 1},
		{maximumNumber: math.MinInt},
	}

	for _, tc := range tests {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("expected panic")
				}
			}()
			var inc = NewIncrementor()
			inc.SetMaximumValue(tc.maximumNumber)
		}()
	}
}

// TestIncrementor_GetNumberAfterSetMaximumValue checks Incrementor.SetMaximumValue
// sets current value to maximum value when it was greater than given maximum value.
func TestIncrementor_GetNumberAfterSetMaximumValue(t *testing.T) {
	const maximumNumber = 5
	var tests = []struct {
		increments int
		want       int
	}{
		{increments: 4, want: 4},
		{increments: 5, want: 5},
		{increments: 6, want: 5},
		{increments: 7, want: 5},
	}

	for _, tc := range tests {
		var inc = NewIncrementor()

		for i := 0; i < tc.increments; i++ {
			inc.IncrementNumber()
		}

		inc.SetMaximumValue(maximumNumber)
		var got = inc.GetNumber()
		if got != tc.want {
			t.Errorf("want %d, got %d", tc.want, got)
		}
	}
}
