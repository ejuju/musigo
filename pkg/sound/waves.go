package sound

import (
	"errors"
	"math"
	"time"
)

// Wave are used for oscillations and effects.
// They produce a value at a certain point in time for a certain float64.
//
// The 1st argument represents the time elapsed when your composition plays throughout time.
// the 1st return value is an output pulse that should be within -1 to 1, it corresponds to a PCM frame value.
// the 2nd return value reports a possible error.
type Wave interface {
	Value(freq float64, elapsed time.Duration) (float64, error)
}

// ErrEndOfWave means that the wave has ended.
var ErrEndOfWave = errors.New("end of wave")

// Sine wave produces an oscillation. It is an implementation of the Wave type.
type SineWave struct{}

func (o *SineWave) Value(freq float64, x time.Duration) (float64, error) {
	return math.Sin(x.Seconds() * 2 * math.Pi * float64(freq)), nil
}

// Square wave produces an oscillation. It is an implementation of the Wave type.
type SquareWave struct{}

func (o *SquareWave) Value(freq float64, x time.Duration) (float64, error) {
	if math.Mod(x.Seconds()*float64(freq), 1) < 0.5 {
		return -1, nil
	}
	return 1, nil
}

// Sawtooth wave produces an oscillation. It is an implementation of the Wave type.
type SawToothWave struct{}

func (o *SawToothWave) Value(freq float64, x time.Duration) (float64, error) {
	return math.Mod(2*x.Seconds()*float64(freq), 2) - 1, nil
}

// WaveWithMaxDuration provides the value of the underlying wave until
// the given duration is reached, after which it returns an ErrEndOfWave.
// The underlying wave can still return an ErrEndOfWave before the deadline is reached.
type WaveWithMaxDuration struct {
	wave     Wave
	duration time.Duration
}

// NewWaveWithMaxDuration creates a new wave with a predefined maximum duration.
// The underlying wave will stop producing a signal when the provided time duration has been reached.
func NewWaveWithMaxDuration(wave Wave, duration time.Duration) *WaveWithMaxDuration {
	return &WaveWithMaxDuration{wave: wave, duration: duration}
}

func (w *WaveWithMaxDuration) Value(freq float64, elapsed time.Duration) (float64, error) {
	if elapsed >= w.duration {
		return 0.0, ErrEndOfWave
	}

	return w.wave.Value(freq, elapsed)
}
