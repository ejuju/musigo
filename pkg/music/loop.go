package music

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

type Loop struct {
	wave         sound.Wave
	iterDuration time.Duration
}

// NewLoop creates a new loop.
// It play the provided wave for the provided duration at each iteration.
// If the underlying wave ends before the end of the duration,
// the wave end is ignored in order to keep the loop going.
func NewLoop(wave sound.Wave, iteration time.Duration) *Loop {
	return &Loop{wave: wave, iterDuration: iteration}
}

func (w *Loop) Value(freq float64, at time.Duration) (float64, error) {
	// override current duration
	at = time.Duration(math.Mod(float64(at), float64(w.iterDuration)))

	val, err := w.wave.Value(freq, at)
	if err != nil {
		// ignore end of wave for the underlying wave
		if errors.Is(err, sound.ErrEndOfWave) {
			return 0.0, nil
		}

		return val, fmt.Errorf("failed to get wave value: %w", err)
	}

	return val, nil
}
