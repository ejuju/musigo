package main

import (
	"time"

	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/music"
	"github.com/ejuju/musigo/pkg/sound"
)

func main() {
	// How to control amplitudes? You use sound.NewWaveWithAmplitudeEnvelope.

	// This example shows how you can make a simple ADSR envelope.
	// This is done using a control wave that will control the amplitude of another wave.
	wave := sound.NewAmplitudeEnvelope(
		sound.NewSynthWave(&sound.Sine{}, music.NoteC5.Hz()),
		sound.NewControlWave(nil, false, []*sound.ControlWaveSegment{
			{Duration: 800 * time.Millisecond, EndValue: 1.0}, // attack
			{Duration: 500 * time.Millisecond, EndValue: 0.7}, // decay
			{Duration: 200 * time.Millisecond, EndValue: 0.2}, // sustain
			{Duration: 2 * time.Second, EndValue: 0.0},        // release
		}),
	)

	err := audio.FFPlayPlayer{Wave: wave, SampleRate: 44100}.Play()
	if err != nil {
		panic(err)
	}
}
