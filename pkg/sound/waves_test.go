package sound

import (
	"testing"
	"time"
)

func TestWithAmplitude(t *testing.T) {
	t.Parallel()

	t.Run("Should apply desired amplitude envelope to a given wave", func(t *testing.T) {
		segments := []AmplitudeEnvelopeSegment{
			{Duration: time.Second, EndValue: 1.0},
			{Duration: time.Second, EndValue: 1.0},
			{Duration: time.Second, EndValue: 0.4},
			{Duration: time.Second, EndValue: 0.0},
		}

		testwave := WithAmplitude(nil, 0, segments...).Wrap(&MockWave{})

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
