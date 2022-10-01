package music

import (
	"math"
	"testing"
)

func TestRelativeFrequency(t *testing.T) {
	tests := []struct {
		root      float64
		semitones float64
		want      float64
	}{
		{root: C3.Frequency(), semitones: 0, want: C3.Frequency()},
		{root: C3.Frequency(), semitones: 2, want: D3.Frequency()},
		{root: C3.Frequency(), semitones: -5, want: G2.Frequency()},
		{root: C3.Frequency(), semitones: 12, want: C4.Frequency()},
		{root: C3.Frequency(), semitones: 14, want: D4.Frequency()},
	}

	for i, test := range tests {
		got := RelativeFrequency(test.root, test.semitones)
		if math.Abs(got-test.want) > 0.01 {
			t.Fatalf("Test %d: want %f, got %f", i, test.want, got)
		}
	}
}

func TestNoteNumbers(t *testing.T) {
	t.Run("Should be midi notation compliant frequencies", func(t *testing.T) {
		tests := []struct {
			note     NoteNumber
			wantFreq float64
		}{
			{note: A0, wantFreq: 27.50},    // lowest midi note
			{note: A4, wantFreq: 440.00},   // standard pitch
			{note: G9, wantFreq: 12543.85}, // highest midi note
		}

		for _, test := range tests {
			got := float64(test.note.Frequency())
			if math.Abs(got-test.wantFreq) > 0.01 {
				t.Fatalf("For note: %d, got: %f but wanted: %f", test.note, got, test.wantFreq)
			}
		}
	})
}
