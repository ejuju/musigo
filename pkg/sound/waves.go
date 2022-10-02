package sound

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/ejuju/musigo/pkg/maths"
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

func (w *MockWave) Value(at time.Duration) (float64, error) {
	return 1.0, nil
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

func (w *WaveWithMaxDuration) Value(elapsed time.Duration) (float64, error) {
	if elapsed >= w.duration {
		return 0.0, ErrEndOfWave
	}

	return w.wave.Value(elapsed)
}

// A control wave is a wave that produces an output meant to be used as a variable by another another wave.
// It can be used to make ADSR envelopes and change the frequency of an oscillator over time.
type ControlWave struct {
	fn       maths.InterpolationFunction
	segments []*ControlWaveSegment
	loop     bool
}

// A control wave segment represents one part of a control wave.
type ControlWaveSegment struct {
	Duration   time.Duration // 1 is one second from previous point (or from 0 if attack)
	StartValue float64
	EndValue   float64
}

// NewControlWave creates a new control wave.
// The linear interpollation function is used as the fallback interpollation function if none is provided.
func NewControlWave(fn maths.InterpolationFunction, loop bool, segments []*ControlWaveSegment) *ControlWave {
	if fn == nil {
		fn = maths.LinearInterpolation{}
	}

	return &ControlWave{fn: fn, loop: loop, segments: segments}
}

// Control waves ignore the provided frequency.
// That's why they are called control waves. They take control.
func (w *ControlWave) Value(at time.Duration) (float64, error) {
	if w.loop {
		at = time.Duration(math.Mod(float64(at), float64(w.Duration())))
	}

	elapsed := time.Duration(0)
	startValue := 0.0

	for _, segment := range w.segments {
		// continue only if x is not in the current segment
		if !(at >= elapsed && at < elapsed+segment.Duration) {
			elapsed += segment.Duration
			startValue = segment.EndValue
			continue
		}

		// if segment start value is defined, set start value to it.
		if math.Abs(segment.StartValue) > 0.000001 {
			startValue = segment.StartValue
		}

		return w.fn.At(
			float64(at),
			float64(elapsed),
			float64(elapsed+segment.Duration),
			startValue,
			segment.EndValue,
		), nil
	}

	// no segment matched, wave ended
	return 0, ErrEndOfWave
}

func (w *ControlWave) Duration() time.Duration {
	totalDur := time.Duration(0)
	for _, segment := range w.segments {
		totalDur += segment.Duration
	}
	return totalDur
}

// AmplitudeEnvelope controls the amplitude of a wave over time.
// This is usually used to make ADSR envelopes.
type AmplitudeEnvelope struct {
	wave        Wave
	controlWave *ControlWave
}

func NewAmplitudeEnvelope(wave Wave, controlWave *ControlWave) *AmplitudeEnvelope {
	return &AmplitudeEnvelope{
		wave:        wave,
		controlWave: controlWave,
	}
}

func (w *AmplitudeEnvelope) Value(x time.Duration) (float64, error) {
	amplitude, err := w.controlWave.Value(x)
	if err != nil {
		return 0.0, fmt.Errorf("unable to get value from control wave: %w", err)
	}

	val, err := w.wave.Value(x)
	if err != nil {
		return 0.0, fmt.Errorf("unable to get value from wave: %w", err)
	}

	return val * amplitude, nil
}

// MergedWaves combines several waves into one.
// It's basically used to make several waves play at the same time.
type MergedWaves struct {
	waves []Wave
}

func NewMergedWaves(waves ...Wave) *MergedWaves {
	return &MergedWaves{waves: waves}
}

func (w *MergedWaves) Value(freq float64, x time.Duration) (float64, error) {
	out := 0.0
	for _, wave := range w.waves {
		val, err := wave.Value(x)
		if err != nil {
			return 0.0, err
		}
		out += val
	}
	return out / float64(len(w.waves)), nil
}
