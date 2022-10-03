package music

import (
	"errors"
	"testing"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

func TestLoop(t *testing.T) {
	t.Parallel()

	t.Run("Should implement sound.Wave interface", func(t *testing.T) {
		var _ sound.Wave = &Loop{}
	})

	t.Run("Should ignore wrapped wave's end", func(t *testing.T) {
		wrapped := sound.NewWaveWithMaxDuration(&sound.MockWave{}, time.Millisecond)
		testwave := NewLoop(wrapped, time.Second)
		if _, err := testwave.Value(2 * time.Millisecond); errors.Is(err, sound.ErrEndOfWave) {
			t.Fatal("shouldn't get error end of wave yet")
		}
	})

	t.Run("Should repeat wrapped wave with the right iteration duration", func(t *testing.T) {
		iterDuration := time.Millisecond

		wrapped := sound.NewWaveWithMaxDuration(&sound.MockWave{}, iterDuration/2)
		testwave := NewLoop(wrapped, iterDuration)

		if got, _ := testwave.Value(100*iterDuration - 1); got != 0.0 {
			t.Fatalf("want 0.0 but got %f", got)
		}
		if got, _ := testwave.Value(100 * iterDuration); got != 1.0 {
			t.Fatalf("want 1.0 but got %f", got)
		}
	})

	t.Run("Should restitute the wrapped wave's value without modification", func(t *testing.T) {
		wrapped := sound.NewWaveWithMaxDuration(&sound.MockWave{}, time.Millisecond)
		testwave := NewLoop(wrapped, time.Millisecond)
		checkAt := time.Millisecond - 1
		got, _ := testwave.Value(checkAt)
		want, _ := wrapped.Value(checkAt)
		if got != want {
			t.Fatalf("want %f but got %f", want, got)
		}
	})
}
