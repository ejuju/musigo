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
	osc := &sound.SineWave{}
	wave := sound.NewWaveWithMaxDuration(osc, time.Second)

	// Play wave with ffplay.
	// (make sure ffplay is installed to play this example)
	// You can try changing the frequency.
	err := audio.FFPlayPlayer{Wave: wave, SampleRate: 44100, Freq: 440.00, SaveFile: true}.Play()
	if err != nil {
		panic(err)
	}
}
