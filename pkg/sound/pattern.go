package sound

import (
	"time"
)

// Pattern represents a wave output bound to time.
// It is used to create sequences of notes, chords and percussions, etc.
//
// You must call NewPattern to create a pattern.
type Pattern struct {
	segments []*PatternSegment
}

// NewPattern instanciates a pattern.
func NewPattern(segments []*PatternSegment) *Pattern {
	return &Pattern{segments: segments}
}

// PatternSegments represents a part of a pattern.
type PatternSegment struct {
	Duration time.Duration
	Wave     Wave
}

func (p *Pattern) Value(x time.Duration) (float64, error) {
	countDuration := time.Duration(0.0)
	for _, segment := range p.segments {
		if !(x >= countDuration && x < countDuration+segment.Duration) {
			countDuration += segment.Duration
			continue
		}

		// don't do anything if wave is not defined,
		// so that segments can be used to represent "silence"
		if segment.Wave == nil {
			return 0, nil
		}

		val, _ := segment.Wave.Value(x - countDuration)
		return val, nil // always return true to enable patterns to go to next segment
	}

	return 0, ErrEndOfWave
}

// Returns the total duration of the pattern segments
func (p *Pattern) Duration() time.Duration {
	out := time.Duration(0)
	for _, segment := range p.segments {
		out += segment.Duration
	}
	return out
}

// Repeat the pattern a number of times.
// If you pass a value of 0, the output pattern will be empty.
// If you pass a value of 1, the output pattern will stay the same.
// If you pass a value of 2, the output pattern is 2 times the input.
func (p *Pattern) Repeat(times int) *Pattern {
	segments := []*PatternSegment{}
	for i := 0; i < times; i++ {
		segments = append(segments, p.segments...)
	}
	p.segments = segments
	return p
}
