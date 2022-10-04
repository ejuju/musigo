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
	wave       Wave
	fn         maths.InterpolationFunction
	startValue float64
	segments   []AmplitudeEnvelopeSegment
}

func WithAmplitude(fn maths.InterpolationFunction, startValue float64, segments ...AmplitudeEnvelopeSegment) AmplitudeEnvelope {
	if fn == nil {
		fn = maths.LinearInterpolation{}
	}
	if len(segments) == 0 {
		segments = []AmplitudeEnvelopeSegment{}
	}
	return AmplitudeEnvelope{fn: fn, startValue: startValue, segments: segments}
}

func (w AmplitudeEnvelope) Append(duration time.Duration, endValue float64) AmplitudeEnvelope {
	w.segments = append(w.segments, AmplitudeEnvelopeSegment{Duration: duration, EndValue: endValue})
	return w
}

func (w AmplitudeEnvelope) Duration() time.Duration {
	totalDur := time.Duration(0)
	for _, segment := range w.segments {
		totalDur += segment.Duration
	}
	return totalDur
}

func (w AmplitudeEnvelope) Wrap(wave Wave) Wave {
	w.wave = wave
	return w
}

func (w AmplitudeEnvelope) Value(at time.Duration) (float64, error) {
	at = time.Duration(math.Mod(float64(at), float64(w.Duration())))
	elapsed := time.Duration(0)
	startValue := w.startValue

	ampl := 0.0
	for _, segment := range w.segments {
		// set val if current time is in this segment
		if at >= elapsed && at < elapsed+segment.Duration {
			ampl = w.fn.At(
				float64(at),
				float64(elapsed),
				float64(elapsed+segment.Duration),
				startValue,
				segment.EndValue,
			)
			break
		}

		elapsed += segment.Duration
		startValue = segment.EndValue
	}

	val, err := w.wave.Value(at)
	if err != nil {
		return 0.0, fmt.Errorf("unable to get value from wave: %w", err)
	}

	return val * ampl, nil
}

// // A control wave is a wave that produces an output meant to be used as a variable by another another wave.
// // It can be used to make ADSR envelopes and change the frequency of an oscillator over time.
// type ControlWave struct {
// 	fn         maths.InterpolationFunction
// 	startValue float64
// 	segments   []*ControlWaveSegment
// }

// A amplitude envelope segment represents one part of a control wave.
type AmplitudeEnvelopeSegment struct {
	Duration time.Duration // 1 is one second from previous point (or from 0 if attack)
	EndValue float64
}

// // NewControlWave creates a new control wave.
// // The linear interpollation function is used as the fallback interpollation function if none is provided.
// func NewControlWave(fn maths.InterpolationFunction, startValue float64, segments []*ControlWaveSegment) *ControlWave {
// 	if fn == nil {
// 		fn = maths.LinearInterpolation{}
// 	}

// 	return &ControlWave{fn: fn, segments: segments, startValue: startValue}
// }

// // Control waves ignore the provided frequency.
// // That's why they are called control waves. They take control.
// func (w *ControlWave) Value(at time.Duration) (float64, error) {
// 	at = time.Duration(math.Mod(float64(at), float64(w.Duration())))
// 	elapsed := time.Duration(0)
// 	startValue := w.startValue

// 	for _, segment := range w.segments {
// 		// continue only if x is not in the current segment
// 		if !(at >= elapsed && at < elapsed+segment.Duration) {
// 			elapsed += segment.Duration
// 			startValue = segment.EndValue
// 			continue
// 		}

// 		return w.fn.At(
// 			float64(at),
// 			float64(elapsed),
// 			float64(elapsed+segment.Duration),
// 			startValue,
// 			segment.EndValue,
// 		), nil
// 	}

// 	// no segment matched, wave ended
// 	return 0, ErrEndOfWave
// }

// func (w *ControlWave) Duration() time.Duration {
// 	totalDur := time.Duration(0)
// 	for _, segment := range w.segments {
// 		totalDur += segment.Duration
// 	}
// 	return totalDur
// }

// func NewTremoloEffect(phase time.Duration) Effect {
// 	return NewAmplitudeEnvelope(NewControlWave(nil, 0, []*ControlWaveSegment{
// 		{Duration: phase / 2, EndValue: 1.0},
// 		{Duration: phase / 2},
// 	}))
// }
