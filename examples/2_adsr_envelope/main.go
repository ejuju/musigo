package main

import (
	"time"

	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/sound"
)

func main() {
	// This example shows how you can make a simple ADSR envelope.
	// This is done using a control wave that will control the amplitude of another wave.
	wave := sound.NewWaveWithAmplitudeEnvelope(
		&sound.SineWave{},
		sound.NewControlWave(nil, []*sound.ControlWaveSegment{
			{Duration: 800 * time.Millisecond, EndValue: 1.0}, // attack
			{Duration: 500 * time.Millisecond, EndValue: 0.7}, // decay
			{Duration: 200 * time.Millisecond, EndValue: 0.2}, // sustain
			{Duration: 2 * time.Second, EndValue: 0.0},        // release
		}),
	)

	err := audio.FFPlayPlayer{Wave: wave, SampleRate: 44100, Freq: 440.00}.Play()
	if err != nil {
		panic(err)
	}
}
