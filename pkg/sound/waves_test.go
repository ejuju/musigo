package sound

import (
	"errors"
	"testing"
	"time"

	"github.com/ejuju/musigo/pkg/maths"
)

func TestWaveWithMaxDuration(t *testing.T) {
	t.Parallel()

	t.Run("Should limit the duration of a wave", func(t *testing.T) {
		duration := 500 * time.Millisecond
		wantValueBeforeEnd := 1.0

		wave := NewWaveWithMaxDuration(&MockWave{}, duration)
		gotValBeforeEnd, gotErrBeforeEnd := wave.Value(duration - 1)
		if gotErrBeforeEnd != nil {
			t.Fatal("wave should not return an error")
		}
		if gotValBeforeEnd != wantValueBeforeEnd {
			t.Fatalf("got: %f, want: %f", gotValBeforeEnd, wantValueBeforeEnd)
		}
		_, gotErrAfterEnd := wave.Value(duration + 1)
		if !errors.Is(gotErrAfterEnd, ErrEndOfWave) {
			t.Fatal("wave should return error end of wave")
		}
	})
}

func TestControlWave(t *testing.T) {
	t.Parallel()

	t.Run("Should create a wave going to certain values", func(t *testing.T) {
		controlWaveSegments := []*ControlWaveSegment{
			{Duration: time.Second, EndValue: 1.0},
			{Duration: time.Second, EndValue: 0.3},
			{Duration: time.Second, EndValue: 0.7},
			{Duration: time.Second, EndValue: 0.0},
		}

		ctrlwave := NewControlWave(maths.LinearInterpolation{}, false, controlWaveSegments)
		durCount := time.Duration(0)

		for _, segment := range controlWaveSegments {
			durCount += segment.Duration
			got, _ := ctrlwave.Value(durCount)
			want := segment.EndValue
			if got != want {
				t.Fatalf("Got %f, want %f", got, want)
			}
		}

		if _, err := ctrlwave.Value(ctrlwave.Duration()); !errors.Is(err, ErrEndOfWave) {
			t.Fatal("control wave should have ended")
		}
	})

	t.Run("Should be loopable", func(t *testing.T) {
		ctrlwave := NewControlWave(maths.LinearInterpolation{}, true, []*ControlWaveSegment{
			{Duration: time.Second, EndValue: 1.0},
		})

		_, gotErr := ctrlwave.Value(time.Second * 2)
		if errors.Is(gotErr, ErrEndOfWave) {
			t.Fatalf("Should't get error end of wave")
		}
	})
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

		testwave := NewAmplitudeEnvelope(&MockWave{}, NewControlWave(nil, false, segments))

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
