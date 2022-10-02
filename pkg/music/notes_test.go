package music

import (
	"math"
	"reflect"
	"testing"
)

func TestRelativeFrequency(t *testing.T) {
	tests := []struct {
		root      float64
		semitones float64
		want      float64
	}{
		{root: NoteC3.Hertz(), semitones: 0, want: NoteC3.Hertz()},
		{root: NoteC3.Hertz(), semitones: 2, want: NoteD3.Hertz()},
		{root: NoteC3.Hertz(), semitones: -5, want: NoteG2.Hertz()},
		{root: NoteC3.Hertz(), semitones: 12, want: NoteC4.Hertz()},
		{root: NoteC3.Hertz(), semitones: 14, want: NoteD4.Hertz()},
	}

	for i, test := range tests {
		got := RelativeFrequency(test.root, test.semitones)
		if math.Abs(got-test.want) > 0.01 {
			t.Fatalf("Test %d: want %f, got %f", i, test.want, got)
		}
	}
}

func TestSemitones(t *testing.T) {
	t.Parallel()

	t.Run("Should be able to control notes' octaves", func(t *testing.T) {
		semitones := SemitonesFromRoot{0, 5, 13, -14}
		octaves := []float64{-1, 0, 1, 2}
		wantSemitones := SemitonesFromRoot{
			-12, -7, 1, -26,
			0, 5, 13, -14,
			12, 17, 25, -2,
			24, 29, 37, 10,
		}

		semitones = semitones.Octave(octaves...)
		for i, wantSemitone := range wantSemitones {
			got := semitones[i]
			if math.Abs(got-wantSemitone) > 0.00001 {
				t.Fatalf("want %f but got %f", wantSemitone, got)
			}
		}
	})

	t.Run("Should be able to be shuffled", func(t *testing.T) {
		semitones := SemitonesFromRoot{1.0, 2.0, 3.0}
		got := make(SemitonesFromRoot, len(semitones))
		copy(got, semitones)
		got = got.Shuffle(0)
		if reflect.DeepEqual(got, semitones) {
			t.Fatalf("Should not be the same as before, got %v but already had %v before", got, semitones)
		}
	})

	t.Run("Should be able to be reversed", func(t *testing.T) {
		notes := SemitonesFromRoot{1.0, 2.0, 3.0}
		want := SemitonesFromRoot{3.0, 2.0, 1.0}
		got := notes.Reverse()
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Should have been reversed, got %#v but want %#v", got, want)
		}
	})
}
