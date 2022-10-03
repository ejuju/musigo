package musigo

import (
	"fmt"
	"time"

	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/sound"
)

type Track func(*TrackController) (error)

type Tracks map[string]func(*TrackController) error

func (w Tracks) Wave(sampleRate int) (sound.Wave, error) {
	// get pulses from all tracks
	allframes := [][]float64{}
	for i, track := range w {
		// create controller
		controller := &TrackController{
			frames:           []float64{},
			durationPerFrame: time.Duration(float64(time.Second) / float64(sampleRate)),
			Synth:            &sound.Sine{},
		}

		// pass controller to track
		err := track(controller)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve frames for current track (at index %s): %w", i, err)
		}

		// get frames from controller
		if len(controller.frames) == 0 {
			return nil, fmt.Errorf("failed to get frames from controller %s", i)
		}

		allframes = append(allframes, controller.frames)
	}

	// calc average val for all frames
	out := []float64{}
	for i := 0; i < len(allframes[0]); i++ {
		total := 0.0
		for _, frames := range allframes {
			total += frames[i]
		}
		out = append(out, total/float64(len(allframes)))
	}
	
	// make a "pattern" of waves and duration instead of calculating frames already
	audio.NewSample(out, sampleRate)

	// apply effects
	for _, v := range  {
		
	}

	return , nil
}

type TrackController struct {
	frames           []float64
	durationPerFrame time.Duration
	Synth            sound.Synthesizer
	Effects          []sound.Effect
	// sample           audio.Sample
}

// Sets the current sample to be used by the controller.
// Can not be used together with SetSynth.
// func (c *TrackController) SetSample(sample audio.Sample) {
// 	c.sample = sample
// }

// Play a sound made of one or more frequencies
func (c *TrackController) Play(duration time.Duration, freqs ...float64) {
	for i := 0; i < int(float64(duration)/float64(c.durationPerFrame)); i++ {
		total := 0.0
		for _, freq := range freqs {
			val := c.Synth.Synthesize(freq, time.Duration(i)*time.Duration(c.durationPerFrame))
			total += val
		}
		sum := total / float64(len(freqs))
		c.frames = append(c.frames, sum)
	}
}

// Like time.Sleep
func (c *TrackController) Sleep(duration time.Duration) {
	for i := 0; i < int(float64(duration)/float64(c.durationPerFrame)); i++ {
		c.frames = append(c.frames, 0.0)
	}
}

//
func (c *TrackController) SleepTillEnd() {

}



func (c *TrackController) PlayTrack(trackID string) {

}
