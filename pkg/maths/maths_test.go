package maths

import (
	"fmt"
	"math"
	"testing"
)

func TestInterpolationFunctions(t *testing.T) {
	t.Parallel()

	funcs := []InterpolationFunction{
		LinearInterpolation{},
	}

	// check if the start value and end values are what they should be.
	// center values can't be checked here and must be checked individually
	tests := []struct {
		x, x1, x2, y1, y2, want float64
	}{
		// going "up"
		{x: 0.0, x1: 0.0, x2: 1.0, y1: 0.0, y2: 1.0, want: 0}, // start (x = 0)
		{x: 1.0, x1: 0.0, x2: 1.0, y1: 0.0, y2: 1.0, want: 1}, // end (x = 1)
		// going "down"
		{x: 0.0, x1: 0.0, x2: 1.0, y1: 1.0, y2: 0.0, want: 1},
		{x: 1.0, x1: 0.0, x2: 1.0, y1: 1.0, y2: 0.0, want: 0},
		// staying "straight"
		{x: 0.0, x1: 0.0, x2: 1.0, y1: 1.0, y2: 1.0, want: 1},
		{x: 1.0, x1: 0.0, x2: 1.0, y1: 1.0, y2: 1.0, want: 1},
	}

	for i, test := range tests {
		for j, fn := range funcs {
			t.Run(fmt.Sprintf("Should return the right value for function %d (test %d)", j, i), func(t *testing.T) {
				got := fn.At(test.x, test.x1, test.x2, test.y1, test.y2)
				if math.Abs(got-test.want) > 0.0000001 {
					t.Fatalf("want %v, got %v", test.want, got)
				}
			})
		}
	}
}
