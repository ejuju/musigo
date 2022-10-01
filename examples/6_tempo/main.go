package main

import (
	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/music"
	"github.com/ejuju/musigo/pkg/sound"
)

func main() {
	// Now, let's say we want to make a loop where the duration of each note lasts one beat.
	// In order to do so, we can calculate a time.Duration given a BPM and a number of beats.
	// Conveniently, the music pkg offers a way to do that.
	// Here, we use it to make a loop.
	bpm := music.BPM(120.0)

	// First, let's make define a wave that we will use for each note
	notewave := sound.NewWaveWithAmplitudeEnvelope(&sound.SineWave{},
		sound.NewControlWave(nil, []*sound.ControlWaveSegment{
			{Duration: bpm.Beats(0.125), EndValue: 1.0},
			{Duration: bpm.Beats(0.125)},
		}),
	)
	loop := music.NewLoop(notewave, bpm.Beats(0.25))

	// Now we want to change the frequency of this loop over time.
	// Let's make a frequency envelope.
	wave := sound.NewWaveWithFrequencyEnvelope(
		loop,
		sound.NewControlWave(nil, []*sound.ControlWaveSegment{
			{Duration: bpm.Beats(4), StartValue: 1.00, EndValue: 1.00},
			{Duration: bpm.Beats(4), StartValue: 1.00, EndValue: 2.50},
			{Duration: bpm.Beats(4), EndValue: 0.50},
			{Duration: bpm.Beats(4), StartValue: 2.00, EndValue: 2.00},
		}),
	)

	// Now we play the result, if we want to change the root frequency, you can do so below.
	err := audio.FFPlayPlayer{Wave: wave, SampleRate: 44100, Freq: 440.00}.Play()
	if err != nil {
		panic(err)
	}
}
