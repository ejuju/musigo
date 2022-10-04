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

	wave := sound.NewWaveWithMaxDuration(
		sound.NewSynthWave(&sound.Sine{}, music.NoteA4.Hz()),
		bpm.Beats(4), // You can call bpm.Beats() to convert any number of beats into a duration.
	)

	// Now we play the result, if we want to change the root frequency, you can do so below.
	err := audio.FFPlayPlayer{Wave: wave, SampleRate: 44100}.Play()
	if err != nil {
		panic(err)
	}
}
