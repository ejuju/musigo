package sound

import (
	"testing"
	"time"
)

func TestPattern(t *testing.T) {
	t.Parallel()

	t.Run("Should create a pattern of waves and durations", func(t *testing.T) {
		osc1 := NewSynthWave(&Sine{}, 440.0)
		osc2 := NewSynthWave(&Square{}, 440.0)
		osc3 := NewSynthWave(&SawTooth{}, 440.0)
		oscillationsInSegmentOrder := []Wave{osc1, osc2, osc3}
		segments := []PatternSegment{
			{Duration: time.Second, Wave: osc1},
			{Duration: time.Second, Wave: osc2},
			{Duration: time.Second, Wave: osc3},
		}

		pattern := NewPattern(segments)

		durCount := time.Duration(0)
		for i, segment := range segments {
			got, _ := pattern.Value(durCount)
			want, _ := oscillationsInSegmentOrder[i].Value(durCount)
			if got != want {
				t.Fatalf("Got %f, want %f", got, want)
			}
			durCount += segment.Duration
		}
	})

	t.Run("Should return the right duration", func(t *testing.T) {
		pattern := NewPattern([]PatternSegment{
			{Duration: 7 * time.Second},
			{Duration: 4 * time.Second},
		})

		got := pattern.Duration()
		want := 11 * time.Second
		if got != want {
			t.Fatalf("Got %s, want %s", got, want)
		}
	})

	t.Run("Should be able to repeat itself", func(t *testing.T) {
		pattern := NewPattern([]PatternSegment{
			{Duration: 1 * time.Second},
		})

		got := pattern.Repeat(2).Duration()
		want := 2 * time.Second
		if got != want {
			t.Fatalf("Got %s, want %s", got, want)
		}
	})
}
