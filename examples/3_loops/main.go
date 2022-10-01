package main

import (
	"time"

	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/music"
	"github.com/ejuju/musigo/pkg/sound"
)

func main() {
	// Let's make a basic sound that we want to loop.
	wave := sound.NewWaveWithAmplitudeEnvelope(
		&sound.SineWave{},
		sound.NewControlWave(nil, []*sound.ControlWaveSegment{
			{Duration: 200 * time.Millisecond, EndValue: 1.0},
			{Duration: 300 * time.Millisecond},
		}),
	)

	// This loop with play the wave, for 500 milliseconds, 16 times.
	loop := music.NewLoop(wave, 500*time.Millisecond, 16)

	err := audio.FFPlayPlayer{Wave: loop, SampleRate: 44100, Freq: 440.00}.Play()
	if err != nil {
		panic(err)
	}
}
