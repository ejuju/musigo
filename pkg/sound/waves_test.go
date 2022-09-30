package sound

import (
	"errors"
	"fmt"
	"math"
	"testing"
	"time"
)

// TestOscillators test if waves oscillate at the right float64 given a certain time.
func TestOscillators(t *testing.T) {
	t.Parallel()

	tests := []struct {
		osc     Wave
		oscName string
	}{
		{oscName: "Sine", osc: &SineWave{}},
		{oscName: "Square", osc: &SquareWave{}},
		{oscName: "Sawtooth", osc: &SawToothWave{}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s should return the same value at two points separated by a \"cycle\"", test.oscName), func(t *testing.T) {
			for i := -1; i < 3; i++ {
				val0, _ := test.osc.Value(float64(i), 0)                            // value at start for a float64 of 1 hertz
				val1, _ := test.osc.Value(float64(i), time.Second*time.Duration(i)) // value after 1 second for a float64 of 1 hertz

				precisionWithinRange := 0.001
				if math.Abs(val0-val1) > precisionWithinRange {
					t.Fatalf("Got: %.3f %.3f", val0, val1)
				}
			}
		})
	}
}

func TestSine(t *testing.T) {
	t.Parallel()

	tests := []struct {
		freq float64
		at   time.Duration
		want float64
	}{
		{freq: 1, at: 0, want: 0.0},                      // start
		{freq: 1, at: 250 * time.Millisecond, want: 1},   // 25% cycle
		{freq: 1, at: 500 * time.Millisecond, want: 0.0}, // mid cycle
		{freq: 1, at: 750 * time.Millisecond, want: -1},  // 75% cycle
		{freq: 1, at: time.Second, want: 0.0},            // full cycle
	}

	for _, test := range tests {
		got, _ := (&SineWave{}).Value(test.freq, test.at)
		if math.Abs(got-test.want) > 0.0001 {
			t.Fatalf("Unexpected value, got %f but want %f", got, test.want)
		}
	}
}

func TestSquare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		freq float64
		at   time.Duration
		want float64
	}{
		{freq: 1, at: 0, want: -1.0},                        // start
		{freq: 1, at: 500*time.Millisecond - 1, want: -1.0}, // before mid cycle
		{freq: 1, at: 500 * time.Millisecond, want: 1.0},    // mid cycle
		{freq: 1, at: time.Second - 1, want: 1.0},           // before end
		{freq: 1, at: time.Second, want: -1.0},              // full cycle
	}

	for _, test := range tests {
		got, _ := (&SquareWave{}).Value(test.freq, test.at)
		if math.Abs(got-test.want) > 0.0001 {
			t.Fatalf("Unexpected value, got %f but want %f", got, test.want)
		}
	}
}

func TestSawTooth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		freq float64
		at   time.Duration
		want float64
	}{
		{freq: 1, at: 0, want: -1.0},                     // start
		{freq: 1, at: 500 * time.Millisecond, want: 0.0}, // mid cycle
		{freq: 1, at: time.Second - 1, want: 1.0},        // end cycle
		{freq: 1, at: time.Second, want: -1.0},           // full cycle
	}

	for _, test := range tests {
		got, _ := (&SawToothWave{}).Value(test.freq, test.at)
		if math.Abs(got-test.want) > 0.0001 {
			t.Fatalf("Unexpected value, got %f but want %f", got, test.want)
		}
	}
}

func TestWaveWithMaxDuration(t *testing.T) {
	t.Parallel()

	t.Run("Should limit the duration of a wave", func(t *testing.T) {
		duration := 500 * time.Millisecond
		wantValueBeforeEnd := -1.0

		wave := NewWaveWithMaxDuration(&SquareWave{}, duration)
		gotValBeforeEnd, gotErrBeforeEnd := wave.Value(1, duration-1)
		if gotErrBeforeEnd != nil {
			t.Fatal("wave should not return an error")
		}
		if gotValBeforeEnd != wantValueBeforeEnd {
			t.Fatalf("got: %f, want: %f", gotValBeforeEnd, wantValueBeforeEnd)
		}
		_, gotErrAfterEnd := wave.Value(1, duration+1)
		if !errors.Is(gotErrAfterEnd, ErrEndOfWave) {
			t.Fatal("wave should return error end of wave")
		}
	})
}
