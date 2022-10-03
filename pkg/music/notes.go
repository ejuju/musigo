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

// Semitones represents the semitones "gaps" between notes of a music scale or chord.
type SemitonesFromRoot []float64

// Octave returns notes containing the octaves passed as inputs.
// Pass zero in the inputs if you want to include the original notes in the output.
func (s SemitonesFromRoot) Octave(octaves ...float64) SemitonesFromRoot {
	out := SemitonesFromRoot{}
	for _, oct := range octaves {
		for _, note := range s {
			out = append(out, note+12*oct)
		}
	}
	return out
}

// Shuffle randomly changes the order of notes.
func (s SemitonesFromRoot) Shuffle(randSeed int64) SemitonesFromRoot {
	r := rand.New(rand.NewSource(randSeed))
	r.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}

// Reverse returns notes in the opposite order.
func (n SemitonesFromRoot) Reverse() SemitonesFromRoot {
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return n
}

func (n SemitonesFromRoot) Frequencies(root float64) []float64 {
	notes := []float64{root}
	for _, gap := range n {
		notes = append(notes, RelativeFrequency(root, gap))
	}
	return notes
}

// // Wave converts a semitones to a sound.Wave
// func (s SemitonesFromRoot) Arpeggio(synth sound.Synthesizer, durations ...time.Duration) *Loop {
// 	segments := []*sound.ControlWaveSegment{}

// 	for i, gap := range s {
// 		duration := durations[i%len(durations)]
// 		segments = append(segments, &sound.ControlWaveSegment{
// 			Duration:   duration,
// 			StartValue: RelativeFrequency(1, gap),
// 			EndValue:   RelativeFrequency(1, gap),
// 		})
// 	}

// 	freqCtrlWave := sound.NewControlWave(nil, segments)
// 	totalDur := freqCtrlWave.Duration()
// 	return NewLoop(sound.NewFrequencyEnvelope(synth, freqCtrlWave), totalDur)
// }

// //
// func (s SemitonesFromRoot) Chord(wave sound.Wave) *sound.MergedWaves {
// 	waves := []sound.Wave{}

// 	for _, gap := range s {
// 		waves = append(waves, sound.NewWaveWithFrequencyMultiplier(wave, RelativeFrequency(1, gap)))
// 	}

// 	return sound.NewMergedWaves(waves...)
// }

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
