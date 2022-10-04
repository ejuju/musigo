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
		{root: NoteC3.Hz(), semitones: 0, want: NoteC3.Hz()},
		{root: NoteC3.Hz(), semitones: 2, want: NoteD3.Hz()},
		{root: NoteC3.Hz(), semitones: -5, want: NoteG2.Hz()},
		{root: NoteC3.Hz(), semitones: 12, want: NoteC4.Hz()},
		{root: NoteC3.Hz(), semitones: 14, want: NoteD4.Hz()},
	}

	for i, test := range tests {
		got := RelativeFrequency(test.root, test.semitones)
		if math.Abs(got-test.want) > 0.01 {
			t.Fatalf("Test %d: want %f, got %f", i, test.want, got)
		}
	}
}

func TestNotes(t *testing.T) {
	t.Parallel()

	t.Run("Should be able to control notes' octaves", func(t *testing.T) {
		notes := Notes{NoteA4.Hz(), NoteC4.Hz(), NoteD2.Hz(), NoteBb6.Hz()}
		octaves := []float64{-1, 0, 1, 2}
		wantSemitones := Notes{
			NoteA3.Hz(), NoteC3.Hz(), NoteD1.Hz(), NoteBb5.Hz(),
			NoteA4.Hz(), NoteC4.Hz(), NoteD2.Hz(), NoteBb6.Hz(),
			NoteA5.Hz(), NoteC5.Hz(), NoteD3.Hz(), NoteBb7.Hz(),
			NoteA6.Hz(), NoteC6.Hz(), NoteD4.Hz(), NoteBb8.Hz(),
		}

		notes = notes.Octave(octaves...)
		for i, wantSemitone := range wantSemitones {
			got := notes[i]
			if math.Abs(got-wantSemitone) > 0.0001 {
				t.Fatalf("want %f but got %f (index: %d)", wantSemitone, got, i)
			}
		}
	})

	t.Run("Should be able to be shuffled", func(t *testing.T) {
		notes := Notes{1.0, 2.0, 3.0}
		got := make(Notes, len(notes))
		copy(got, notes)
		got = got.Shuffle(0)
		if reflect.DeepEqual(got, notes) {
			t.Fatalf("Should not be the same as before, got %v but already had %v before", got, notes)
		}
	})

	t.Run("Should be able to be reversed", func(t *testing.T) {
		notes := Notes{1.0, 2.0, 3.0}
		want := Notes{3.0, 2.0, 1.0}
		got := notes.Reverse()
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Should have been reversed, got %#v but want %#v", got, want)
		}
	})
}
