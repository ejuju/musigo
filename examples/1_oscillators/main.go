package main

import (
	"time"

	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/sound"
)

func main() {
	// Common oscillators (sine, square, etc.) are available.
	// This example plays them one by one to hear what they sound like.
	oscillators := []sound.Wave{
		&sound.SineWave{},
		&sound.SquareWave{},
		&sound.SawToothWave{},
	}

	for _, osc := range oscillators {
		wave := sound.NewWaveWithMaxDuration(osc, time.Second)

		err := audio.FFPlayPlayer{Wave: wave, SampleRate: 44100, Freq: 440.00}.Play()
		if err != nil {
			panic(err)
		}
	}
}
