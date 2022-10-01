package music

import (
	"math"
	"testing"
)

func TestBPM(t *testing.T) {
	t.Parallel()

	t.Run("Should return the right duration", func(t *testing.T) {
		tests := []struct {
			bpm     BPM
			beats   float64
			wantSec float64
		}{
			{bpm: 60, beats: 2.00, wantSec: 2.000},
			{bpm: 60, beats: 1.00, wantSec: 1.000},
			{bpm: 60, beats: 0.25, wantSec: 0.250},
			{bpm: 120, beats: 2.00, wantSec: 1.000},
			{bpm: 120, beats: 1.00, wantSec: 0.500},
			{bpm: 120, beats: 0.25, wantSec: 0.125},
		}

		for _, test := range tests {
			val := test.bpm.Beats(test.beats).Seconds()
			if math.Abs(val-test.wantSec) > 0.001 {
				t.Fatalf("Want: %.5fs, got: %.5fs", test.wantSec, val)
			}
		}
	})
}
