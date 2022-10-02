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

	// Now we want to change the frequency of this loop over time.
	// Let's define our melody.
	wave := sound.NewFrequencyEnvelope(
		&sound.Sine{},
		sound.NewControlWave(nil, false, []*sound.ControlWaveSegment{
			{Duration: noteDuration * 4, StartValue: music.NoteA4.Hertz(), EndValue: music.NoteA4.Hertz()},
			{Duration: noteDuration * 4, StartValue: music.NoteB4.Hertz(), EndValue: music.NoteB4.Hertz()},
			{Duration: noteDuration * 4, StartValue: music.NoteD4.Hertz(), EndValue: music.NoteD4.Hertz()},
			{Duration: noteDuration * 4, StartValue: music.NoteE4.Hertz(), EndValue: music.NoteE4.Hertz()},
		}),
	)

	// First, let's make a basic sound loop with the note.
	out := sound.NewAmplitudeEnvelope(
		wave,
		sound.NewControlWave(nil, true, []*sound.ControlWaveSegment{
			{Duration: 1 * noteDuration / 3, EndValue: 1.0},
			{Duration: 2 * noteDuration / 3},
		}),
	)

	// Now we play the result, if we want to change the root frequency, you can do so below.
	err := audio.FFPlayPlayer{Wave: sound.NewWaveWithMaxDuration(out, noteDuration*32), SampleRate: 44100}.Play()
	if err != nil {
		panic(err)
	}
}
