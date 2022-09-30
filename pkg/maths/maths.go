package maths

import (
	"math"
)

type InterpolationFunction interface {
	At(x, x1, x2, y1, y2 float64) float64
}

type LinearInterpolation struct{}

func (i LinearInterpolation) At(x, x1, x2, y1, y2 float64) float64 {
	deltaX := math.Abs(x2 - x1)
	deltaY := y2 - y1
	progressX := float64(x-x1) / float64(deltaX)
	return progressX*deltaY + y1
}
