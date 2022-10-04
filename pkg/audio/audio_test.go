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
		sampleRate := 2
		// generate frames from a square wave and check if the output frames are what we expect
		wave := sound.NewSynthWave(&sound.Square{}, 1)
		frames, _ := Frames(wave, sampleRate, 0, 2*time.Second)

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
			frames, _ := Frames(&sound.MockWave{}, test.sampleRate, 0, test.duration)
			got := len(frames)
			if got != test.want {
				t.Fatalf("Got: %d, but wanted: %d", got, test.want)
			}
		}
	})

	t.Run("Should validate required inputs", func(t *testing.T) {
		validWave := &sound.MockWave{}
		var invalidWave sound.Wave = nil
		validSampleRate := 1
		invalidSampleRate := 0

		tests := []struct {
			wave       sound.Wave
			sampleRate int
			wantErr    bool
		}{
			{wave: validWave, sampleRate: validSampleRate, wantErr: false},  // all valid
			{wave: invalidWave, sampleRate: validSampleRate, wantErr: true}, // only invalid wave
			{wave: validWave, sampleRate: invalidSampleRate, wantErr: true}, // only invalid sample rate
		}

		for i, test := range tests {
			_, err := Frames(test.wave, test.sampleRate, 0, time.Second)
			if (err != nil) != test.wantErr {
				t.Fatalf("unexpected error at index %d, wantErr is %v but got %v", i, test.wantErr, err)
			}
		}
	})
}
