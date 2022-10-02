package audio

import (
	"errors"
	"testing"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

func TestSample(t *testing.T) {
	t.Parallel()

	t.Run("Should implement sound.Wave", func(t *testing.T) {
		var _ sound.Wave = &Sample{}
	})

	t.Run("Should return the right frame value for a given time", func(t *testing.T) {
		frames := []float64{1.0, 0.0, -1.0}
		sampleRate := 44100
		sample := NewSample(frames, sampleRate)

		for i, frame := range frames {
			elapsed := time.Second * time.Duration(i) / time.Duration(sampleRate)
			got, _ := sample.Value(elapsed)
			if got != frame {
				t.Fatalf("unexpected frame value at index %d, want %f but got %f", i, frame, got)
			}
		}
	})

	t.Run("Should return ErrEndOfWave when there's no frame left", func(t *testing.T) {
		frames := []float64{1.0, 0.0, -1.0}
		sampleRate := 1
		sample := NewSample(frames, sampleRate)

		timeAfterFramesEnd := time.Second * time.Duration(len(frames))
		_, err := sample.Value(timeAfterFramesEnd)
		if !errors.Is(err, sound.ErrEndOfWave) {
			t.Fatalf("unexpected error value, want %s but got %s", sound.ErrEndOfWave, err)
		}
	})
}
