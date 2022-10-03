package musigo

import (
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

// Track is an individual sound used in a mix.
// A track could be a bassline, percussions, or a sample.
type Track struct {
	synth     sound.Synthesizer
	trackFunc TrackFunc
	effects   []sound.Effect
}

// TrackFunc is a callback function that gets called when the track gets played.
type TrackFunc func(*Controller)

// Controller allows users to control a song.
type Controller struct {
	t        *Track
	segments []*sound.PatternSegment
}

// NewTrack creates a new track.
func NewTrack(synth sound.Synthesizer, trackFunc TrackFunc, baseEffects ...sound.Effect) *Track {
	return &Track{synth: synth, trackFunc: trackFunc, effects: baseEffects}
}

// Tracks is a map of track IDs and their corresponding track.
type Tracks map[string]*Track

func (tracks Tracks) Merge() *sound.MergedWaves {
	waves := []sound.Wave{}
	for _, track := range tracks {
		controller := &Controller{segments: []*sound.PatternSegment{}, t: track}
		track.trackFunc(controller)
		waves = append(waves, sound.NewPattern(controller.segments))
	}
	return sound.NewMergedWaves(waves...)
}

// Play a sound made of one or more frequencies
func (c *Controller) Play(duration time.Duration, effects []sound.Effect, freqs ...float64) {
	// merge synth waves frequencies into one wave
	waves := []sound.Wave{}
	for _, freq := range freqs {
		waves = append(waves, sound.NewSynthWave(c.t.synth, freq))
	}
	var wave sound.Wave = sound.NewMergedWaves(waves...)

	// wrap effects around wave
	for _, effect := range append(effects, c.t.effects...) {
		wave = effect.Wrap(wave)
	}

	// add segment to controller
	c.segments = append(c.segments, &sound.PatternSegment{
		Duration: duration,
		Wave:     wave,
	})
}

// Like time.Sleep for your track.
// if duration == 0, then will sleep until the end of the track.
func (c *Controller) Wait(duration time.Duration) {
	// add segment to controller
	c.segments = append(c.segments, &sound.PatternSegment{
		Duration: duration,
		Wave:     &sound.SilentWave{},
	})
}
