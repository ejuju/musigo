package main

import (
	"time"

	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/music"
	"github.com/ejuju/musigo/pkg/sound"
)

func main() {
	// For now, we've been using frequencies instead of notes.
	// Here's how you can use notes.
	frequency := music.NoteC4.Hz()

	// let's make a simple wave that lasts for one second to hear the note.
	osc := sound.NewSynthWave(&sound.Sine{}, frequency)
	wave := sound.NewWaveWithMaxDuration(osc, time.Second)

	err := audio.FFPlayPlayer{Wave: wave, SampleRate: 44100}.Play()
	if err != nil {
		panic(err)
	}
}
