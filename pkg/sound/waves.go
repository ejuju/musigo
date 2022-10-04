package sound

import (
	"errors"
	"time"
)

// Wave are used for oscillations and effects.
// They produce a value at a certain point in time for a certain float64.
//
// The 1st argument represents the time elapsed when your composition plays throughout time.
// the 1st return value is an output pulse that should be within -1 to 1, it corresponds to a PCM frame value.
// the 2nd return value reports a possible error.
type Wave interface {
	Value(elapsed time.Duration) (float64, error)
}

// ErrEndOfWave means that the wave has ended.
var ErrEndOfWave = errors.New("end of wave")

// MockWave is a wave that always produces the value of one.
// Useful for testing and debugging.
type MockWave struct{}

func (w MockWave) Value(at time.Duration) (float64, error) {
	return 1.0, nil
}

// SilentWave implements the Wave interface but produces no sound.
type SilentWave struct{}

func (w SilentWave) Value(at time.Duration) (float64, error) {
	return 0, nil
}

// MergedWaves combines several waves into one.
// It's basically used to make several waves play at the same time.
type MergedWaves struct {
	waves []Wave
}

func NewMergedWaves(waves ...Wave) MergedWaves {
	return MergedWaves{waves: waves}
}

func (w MergedWaves) Value(x time.Duration) (float64, error) {
	out := 0.0
	for _, wave := range w.waves {
		val, err := wave.Value(x)
		if err != nil {
			if errors.Is(err, ErrEndOfWave) {
				out += 0
				continue
			}
			return 0.0, err
		}
		out += val
	}
	return out / float64(len(w.waves)), nil
}
