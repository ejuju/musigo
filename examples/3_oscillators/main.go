package main

import (
	"time"

	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/music"
	"github.com/ejuju/musigo/pkg/sound"
)

func main() {
	// Common oscillators (sine, square, etc.) are available.
	// This example plays them one by one to hear what they sound like.
	oscillators := []sound.Wave{
		sound.NewSynthWave(&sound.Sine{}, music.NoteA4.Hz()),
		sound.NewSynthWave(&sound.Square{}, music.NoteA4.Hz()),
		sound.NewSynthWave(&sound.SawTooth{}, music.NoteA4.Hz()),
	}

	for _, osc := range oscillators {
		wave := sound.NewWaveWithMaxDuration(osc, time.Second)
		err := audio.FFPlayPlayer{Wave: wave, SampleRate: 44100}.Play()
		if err != nil {
			panic(err)
		}
	}
}
