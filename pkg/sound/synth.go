package sound

import (
	"math"
	"math/rand"
	"time"
)

// Synthesizer can synthesize sound with a given frequency.
// It produces an oscillation.
type Synthesizer interface {
	Synthesize(freq float64, at time.Duration) (float64, error)
}

// SynthWave allows conversion from Synthesizer to Wave.
type SynthWave struct {
	synth Synthesizer
	freq  float64
}

// NewSynthWave converts a Synthesizer into a Wave by setting a specific frequency.
func NewSynthWave(synth Synthesizer, frequency float64) SynthWave {
	return SynthWave{synth: synth, freq: frequency}
}

func (w SynthWave) Value(at time.Duration) (float64, error) {
	return w.synth.Synthesize(w.freq, at)
}

// Sine produces an oscillation. It is an implementation of the Synthesizer type.
type Sine struct{}

func (o Sine) Synthesize(freq float64, x time.Duration) (float64, error) {
	return math.Sin(x.Seconds() * 2 * math.Pi * float64(freq)), nil
}

// Square produces an oscillation. It is an implementation of the Synthesizer type.
type Square struct{}

func (o Square) Synthesize(freq float64, x time.Duration) (float64, error) {
	if math.Mod(x.Seconds()*float64(freq), 1) < 0.5 {
		return -1, nil
	}
	return 1, nil
}

// Sawtooth produces an oscillation. It is an implementation of the Synthesizer type.
type SawTooth struct{}

func (o SawTooth) Synthesize(freq float64, x time.Duration) (float64, error) {
	return math.Mod(2*x.Seconds()*float64(freq), 2) - 1, nil
}

// RandomWideBandNoiseSynthesizer returns wide band random noise (mmmkay).
// It is used for substractive synthesis (in combination with filters and envelopes)
type RandomWideBandNoiseSynthesizer struct {
	rand *rand.Rand
}

func NewRandomWideBandNoiseSynthesizer(seed int64) RandomWideBandNoiseSynthesizer {
	return RandomWideBandNoiseSynthesizer{rand: rand.New(rand.NewSource(seed))}
}

func (p RandomWideBandNoiseSynthesizer) Synthesize(freq float64, x time.Duration) (float64, error) {
	return (p.rand.Float64() * 2) - 1.0, nil
}
