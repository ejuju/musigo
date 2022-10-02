package audio

import (
	"fmt"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

// Sample represents an audio sample
// (for ex: a drum sound that you want to use in your song)
type Sample struct {
	name       string
	sampleRate int // how many frames per second
	frames     []float64
}

// NewSample creates a new sample using the input frames and sample rate.
func NewSample(frames []float64, sampleRate int) *Sample {
	return &Sample{frames: frames, sampleRate: sampleRate}
}

// Value returns the frame value for the at the input time duration.
func (s *Sample) Value(at time.Duration) (float64, error) {
	frameIndex := at / (time.Second / time.Duration(s.sampleRate))

	if int(frameIndex) >= len(s.frames) {
		return 0, fmt.Errorf("failed to get frame index for sample \"%s\": %w", s.name, sound.ErrEndOfWave)
	}

	return s.frames[frameIndex], nil
}
