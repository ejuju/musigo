package main

import (
	"time"

	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/music"
	"github.com/ejuju/musigo/pkg/sound"
)

func main() {
	// Chords are defined using semitones intervals.
	chord := music.ChordMajor7.Chord(&sound.SineWave{})
	wave := sound.NewWaveWithMaxDuration(chord, 2*time.Second)

	err := audio.FFPlayPlayer{Freq: music.NoteA4.Frequency(), Wave: wave, SampleRate: 44100}.Play()
	if err != nil {
		panic(err)
	}
}
