package music

import (
	"errors"
	"math"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

type Loop struct {
	wave         sound.Wave
	iterDuration time.Duration
	times        int
}

// NewLoop creates a new loop.
// It play the provided wave N times for the provided duration each time.
// If the underlying wave ends before the end of the duration,
// the wave end is ignored in order to keep the loop going.
func NewLoop(wave sound.Wave, dur time.Duration, times int) *Loop {
	return &Loop{wave: wave, iterDuration: dur, times: times}
}

func (w *Loop) Value(freq float64, at time.Duration) (float64, error) {
	// check if end of wave has been reached
	if at > w.iterDuration*time.Duration(w.times) {
		return 0.0, sound.ErrEndOfWave
	}

	// override current duration
	at = time.Duration(math.Mod(float64(at), float64(w.iterDuration)))

	val, err := w.wave.Value(freq, at)
	// ignore end of wave for the underlying wave
	if errors.Is(err, sound.ErrEndOfWave) {
		return 0.0, nil
	}

	return val, nil
}
