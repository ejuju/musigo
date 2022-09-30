package audio

import (
	"testing"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

func TestFrames(t *testing.T) {
	t.Parallel()

	t.Run("Should generate frames with the right values", func(t *testing.T) {
		wantFrames := []float64{-1.0, 1.0, -1.0, 1.0}
		sampleRate := 44100
		// generate frames from a square wave and check if the output frames are what we expect
		wave := sound.NewWaveWithMaxDuration(&sound.SquareWave{}, time.Duration(len(wantFrames)))
		frames, _ := Frames(1, wave, sampleRate)

		for i, got := range frames {
			if got != wantFrames[i] {
				t.Fatalf("unexpected frame value at index %d, want %f but got %f", i, wantFrames[i], got)
			}
		}
	})

	t.Run("Should generate the right number of frames", func(t *testing.T) {
		tests := []struct {
			sampleRate int
			duration   time.Duration
			want       int // number of output samples
		}{
			{sampleRate: 2, duration: 10 * time.Second, want: 20},
			{sampleRate: 10, duration: time.Second, want: 10},
			{sampleRate: 10, duration: 2 * time.Second, want: 20},
			{sampleRate: 100, duration: 10 * time.Second, want: 1000},
		}

		for _, test := range tests {
			osc := &sound.SineWave{}
			frames, _ := Frames(440, sound.NewWaveWithMaxDuration(osc, test.duration), test.sampleRate)
			got := len(frames)
			if got != test.want {
				t.Fatalf("Got: %d, but wanted: %d", got, test.want)
			}
		}
	})
}
