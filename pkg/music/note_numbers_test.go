package music

import (
	"math"
	"testing"
)

func TestNoteNumbers(t *testing.T) {
	t.Run("Should be midi notation compliant frequencies", func(t *testing.T) {
		tests := []struct {
			note     NoteNumber
			wantFreq float64
		}{
			{note: NoteA0, wantFreq: 27.50},    // lowest midi note
			{note: NoteA4, wantFreq: 440.00},   // standard pitch
			{note: NoteG9, wantFreq: 12543.85}, // highest midi note
		}

		for _, test := range tests {
			got := float64(test.note.Hertz())
			if math.Abs(got-test.wantFreq) > 0.01 {
				t.Fatalf("For note: %d, got: %f but wanted: %f", test.note, got, test.wantFreq)
			}
		}
	})
}
