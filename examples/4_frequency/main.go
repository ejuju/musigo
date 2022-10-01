package main

import (
	"time"

	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/music"
	"github.com/ejuju/musigo/pkg/sound"
)

func main() {
	// How to control amplitudes? You use sound.NewWaveWithFrequencyEnvelope.

	noteDuration := 200 * time.Millisecond

	// First, let's make a basic sound loop with the note.
	note := sound.NewWaveWithAmplitudeEnvelope(
		&sound.SineWave{},
		sound.NewControlWave(nil, []*sound.ControlWaveSegment{
			{Duration: 1 * noteDuration / 3, EndValue: 1.0},
			{Duration: 2 * noteDuration / 3},
		}),
	)
	loop := music.NewLoop(note, noteDuration)

	// Now we want to change the frequency of this loop over time.
	// Let's make a frequency envelope.
	wave := sound.NewWaveWithFrequencyEnvelope(
		loop,
		sound.NewControlWave(nil, []*sound.ControlWaveSegment{
			{Duration: noteDuration * 4, StartValue: 1.00, EndValue: 1.00},
			{Duration: noteDuration * 4, StartValue: 1.00, EndValue: 2.50},
			{Duration: noteDuration * 4, EndValue: 0.50},
			{Duration: noteDuration * 4, StartValue: 2.00, EndValue: 2.00},
		}),
	)

	// Now we play the result, if we want to change the root frequency, you can do so below.
	err := audio.FFPlayPlayer{Wave: wave, SampleRate: 44100, Freq: 440.00}.Play()
	if err != nil {
		panic(err)
	}
}
