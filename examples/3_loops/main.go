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

	// This loop with play the wave on repeat.
	loop := music.NewLoop(wave, 500*time.Millisecond)

	// set max duration otherwise loop goes on forever.
	out := sound.NewWaveWithMaxDuration(loop, 16*time.Second)
	err := audio.FFPlayPlayer{Wave: out, SampleRate: 44100, Freq: 440.00}.Play()
	if err != nil {
		panic(err)
	}
}
