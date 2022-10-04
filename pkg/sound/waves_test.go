package sound

import (
	"testing"
	"time"

	"github.com/ejuju/musigo/pkg/maths"
)

func TestControlWave(t *testing.T) {
	t.Parallel()

	t.Run("Should create a wave going to certain values", func(t *testing.T) {
		controlWaveSegments := []*ControlWaveSegment{
			{Duration: time.Second, EndValue: 1.0},
			{Duration: time.Second, EndValue: 0.3},
			{Duration: time.Second, EndValue: 0.7},
			{Duration: time.Second, EndValue: 0.0},
		}

		ctrlwave := NewControlWave(maths.LinearInterpolation{}, 0, controlWaveSegments)
		durCount := time.Duration(0)

		// outer loop to make sure control wave repeats itself at least twice
		for i := 0; i < 3; i++ {
			// check that the value at each segment end is the right one
			for _, segment := range controlWaveSegments {
				durCount += segment.Duration
				got, _ := ctrlwave.Value(durCount)
				want := segment.EndValue
				if got != want {
					t.Fatalf("Got %f, want %f", got, want)
				}
			}
		}
	})

	// t.Run("Should be loopable", func(t *testing.T) {
	// 	ctrlwave := NewControlWave(maths.LinearInterpolation{}, []*ControlWaveSegment{
	// 		{Duration: time.Second, EndValue: 1.0},
	// 	})

	// 	_, gotErr := ctrlwave.Value(time.Second * 2)
	// 	if errors.Is(gotErr, ErrEndOfWave) {
	// 		t.Fatalf("Should't get error end of wave")
	// 	}
	// })
}

func TestWaveWithAmplitudeEnvelope(t *testing.T) {
	t.Parallel()

	t.Run("Should apply desired amplitude envelope to a given wave", func(t *testing.T) {
		segments := []*ControlWaveSegment{
			{Duration: time.Second, EndValue: 1.0},
			{Duration: time.Second, EndValue: 1.0},
			{Duration: time.Second, EndValue: 0.4},
			{Duration: time.Second, EndValue: 0.0},
		}

		testwave := NewAmplitudeEnvelope(NewControlWave(nil, 0, segments)).Wrap(&MockWave{})

		countDur := time.Duration(0)
		for _, segment := range segments {
			gotAtSegmentEnd, _ := testwave.Value(countDur + segment.Duration)
			countDur += segment.Duration // inc for next iteration

			// check if segment end value is right
			if gotAtSegmentEnd != segment.EndValue {
				t.Fatalf("want %f at end of segment but got %f", segment.EndValue, gotAtSegmentEnd)
			}
		}
	})
}
