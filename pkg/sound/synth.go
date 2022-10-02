package sound

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Synthesizer can synthesize sound with a given frequency.
// It produces an oscillation.
type Synthesizer interface {
	Synthesize(freq float64, at time.Duration) float64
}

// Sine produces an oscillation. It is an implementation of the Synthesizer type.
type Sine struct{}

func (o *Sine) Synthesize(freq float64, x time.Duration) float64 {
	return math.Sin(x.Seconds() * 2 * math.Pi * float64(freq))
}

// Square produces an oscillation. It is an implementation of the Synthesizer type.
type Square struct{}

func (o *Square) Synthesize(freq float64, x time.Duration) float64 {
	if math.Mod(x.Seconds()*float64(freq), 1) < 0.5 {
		return -1
	}
	return 1
}

// Sawtooth produces an oscillation. It is an implementation of the Synthesizer type.
type SawTooth struct{}

func (o *SawTooth) Synthesize(freq float64, x time.Duration) float64 {
	return math.Mod(2*x.Seconds()*float64(freq), 2) - 1
}

// FrequencyEnvelope controls the frequency of a wave over time.
//
// It multiplies the frequency passed to the wave by the value of the control wave.
// The control wave should return 1 to keep the same frequency.
// Numbers above 1 will increase the frequency and below 1 will decrease the frequency.
type FrequencyEnvelope struct {
	synth       Synthesizer
	controlWave *ControlWave
}

func NewFrequencyEnvelope(synth Synthesizer, controlWave *ControlWave) *FrequencyEnvelope {
	return &FrequencyEnvelope{
		synth:       synth,
		controlWave: controlWave,
	}
}

func (w *FrequencyEnvelope) Value(x time.Duration) (float64, error) {
	ctrlfreq, err := w.controlWave.Value(x)
	if err != nil {
		return 0.0, fmt.Errorf("unable to get value from control wave: %w", err)
	}

	return w.synth.Synthesize(ctrlfreq, x), nil
}

// SynthWithFrequencyMultiplier multiplies the frequency by the multiplier and passes it to the underlying synth.
type SynthWithFrequencyMultiplier struct {
	synth      Synthesizer
	multiplier float64
}

func NewSynthWithFrequencyMultiplier(synth Synthesizer, multiplier float64) *SynthWithFrequencyMultiplier {
	return &SynthWithFrequencyMultiplier{synth: synth, multiplier: multiplier}
}

func (w *SynthWithFrequencyMultiplier) Synthesize(freq float64, at time.Duration) float64 {
	return w.synth.Synthesize(freq*w.multiplier, at)
}

// RandomWideBandNoiseSynthesizer returns wide band random noise (mmmkay).
// It is used for substractive synthesis (in combination with filters and envelopes)
type RandomWideBandNoiseSynthesizer struct {
	rand *rand.Rand
}

func NewRandomWideBandNoiseSynthesizer(seed int64) *RandomWideBandNoiseSynthesizer {
	return &RandomWideBandNoiseSynthesizer{rand: rand.New(rand.NewSource(seed))}
}

func (p *RandomWideBandNoiseSynthesizer) Synthesize(freq float64, x time.Duration) float64 {
	return (p.rand.Float64() * 2) - 1.0
}

type SynthWave struct {
	synth Synthesizer
	freq  float64
}

func NewSynthWave(synth Synthesizer, frequency float64) *SynthWave {
	return &SynthWave{synth: synth, freq: frequency}
}

func (w *SynthWave) Value(at time.Duration) (float64, error) {
	return w.synth.Synthesize(w.freq, at), nil
}
