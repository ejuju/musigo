package music

import (
	"math"
	"math/rand"
)

// The constant used in the calculation to define music notes.
var MusicalConstant = math.Pow(2, 1.0/12.0)

// RelativeFrequency calculates the frequency of a note given a root frequency and number of semitones.
func RelativeFrequency(root float64, semitones float64) float64 {
	return float64(root) * math.Pow(MusicalConstant, semitones)
}

type Notes []float64

// Octave returns notes containing the octaves passed as inputs.
// Pass zero in the inputs if you want to include the original notes in the output.
func (n Notes) Octave(octaves ...float64) Notes {
	out := Notes{}
	for _, oct := range octaves {
		for _, note := range n {
			out = append(out, RelativeFrequency(note, 12*oct))
		}
	}
	return out
}

// Shuffle randomly changes the order of notes.
func (n Notes) Shuffle(randSeed int64) Notes {
	r := rand.New(rand.NewSource(randSeed))
	r.Shuffle(len(n), func(i, j int) {
		n[i], n[j] = n[j], n[i]
	})
	return n
}

// Reverse returns notes in the opposite order.
func (n Notes) Reverse() Notes {
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return n
}

// Repeat returns notes in the opposite order.
func (n Notes) Repeat(times int) Notes {
	for i := 0; i < times; i++ {
		n = append(n, n...)
	}
	return n
}

// Semitones represents the semitones "gaps" between notes of a music scale or chord.
type SemitonesFromRoot []float64

func (n SemitonesFromRoot) Notes(root float64) Notes {
	notes := Notes{root}
	for _, gap := range n {
		notes = append(notes, RelativeFrequency(root, gap))
	}
	return notes
}

// All pre-defined chords in one slice.
var AllChords = []SemitonesFromRoot{
	ChordMajor,
	ChordMinor,
	ChordMajor7,
	ChordMinor7,
	Chord5,
}

// Common chords are defined as relative semitones intervals.
var (
	ChordMajor  = SemitonesFromRoot{3, 5}
	ChordMinor  = SemitonesFromRoot{2, 5}
	ChordMajor7 = SemitonesFromRoot{3, 5, 10}
	ChordMinor7 = SemitonesFromRoot{2, 3, 10}
	Chord5      = SemitonesFromRoot{5}
)

// All pre-defined scales in one slice.
var AllScales = []SemitonesFromRoot{
	HarmonicMajorScale,
	HarmonicMinorScale,
	PentatonicMinorScale,
	BluesMinorScale,
}

var (
	HarmonicMajorScale   = SemitonesFromRoot{2, 4, 5, 7, 9, 11}
	HarmonicMinorScale   = SemitonesFromRoot{2, 3, 5, 7, 8, 11}
	PentatonicMinorScale = SemitonesFromRoot{3, 5, 7, 10}
	BluesMinorScale      = SemitonesFromRoot{3, 5, 6, 7, 10}
)
