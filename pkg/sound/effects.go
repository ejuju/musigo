package sound

import (
	"fmt"
	"math"
	"time"

	"github.com/ejuju/musigo/pkg/maths"
)

type Effect interface {
	Wave
	Wrap(Wave) Wave
}

// AmplitudeEnvelope is an effect that controls the amplitude of a wave over time.
// This is usually used to make ADSR envelopes.
type AmplitudeEnvelope struct {
	wave        Wave
	controlWave *ControlWave
}

func NewAmplitudeEnvelope(controlWave *ControlWave) *AmplitudeEnvelope {
	return &AmplitudeEnvelope{controlWave: controlWave}
}

func (w *AmplitudeEnvelope) Wrap(wave Wave) Wave {
	w.wave = wave
	return w
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

// A control wave is a wave that produces an output meant to be used as a variable by another another wave.
// It can be used to make ADSR envelopes and change the frequency of an oscillator over time.
type ControlWave struct {
	fn       maths.InterpolationFunction
	segments []*ControlWaveSegment
}

// A control wave segment represents one part of a control wave.
type ControlWaveSegment struct {
	Duration   time.Duration // 1 is one second from previous point (or from 0 if attack)
	StartValue float64
	EndValue   float64
}

// NewControlWave creates a new control wave.
// The linear interpollation function is used as the fallback interpollation function if none is provided.
func NewControlWave(fn maths.InterpolationFunction, segments []*ControlWaveSegment) *ControlWave {
	if fn == nil {
		fn = maths.LinearInterpolation{}
	}

	return &ControlWave{fn: fn, segments: segments}
}

// Control waves ignore the provided frequency.
// That's why they are called control waves. They take control.
func (w *ControlWave) Value(at time.Duration) (float64, error) {
	at = time.Duration(math.Mod(float64(at), float64(w.Duration())))
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

func NewTremoloEffect(phase time.Duration) Effect {
	return NewAmplitudeEnvelope(NewControlWave(nil, []*ControlWaveSegment{
		{Duration: phase / 2, EndValue: 1.0},
		{Duration: phase / 2},
	}))
}
