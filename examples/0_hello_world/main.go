package main

import (
	"time"

	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/sound"
)

func main() {
	// You will most likely not be using Musigo likes it is in this example.
	// This should at least make sure that you're good to go, meaning you can create
	// sound with Musigo and can move on to further adventures.

	// Create a wave that lasts one second and plays a 440.00 hertz frequency sound using a sine oscillator.
	// You can try changing the oscillator and the duration.
	osc := sound.NewSynthWave(&sound.Sine{}, 440.00)

	// Play wave with ffplay.
	// (make sure ffplay is installed to play this example)
	// You can try changing the frequency.
	err := audio.FFPlayPlayer{Wave: osc, SampleRate: 44100, SaveFile: true, Duration: time.Second}.Play()
	if err != nil {
		panic(err)
	}
}
