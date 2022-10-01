package music

import (
	"errors"
	"testing"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

func TestLoop(t *testing.T) {
	t.Run("Should ignore wrapped wave's end", func(t *testing.T) {
		wrapped := sound.NewWaveWithMaxDuration(&sound.MockWave{}, time.Millisecond)
		testwave := NewLoop(wrapped, time.Second, 2)
		if _, err := testwave.Value(0, 2*time.Millisecond); errors.Is(err, sound.ErrEndOfWave) {
			t.Fatal("shouldn't get error end of wave yet")
		}
	})

	t.Run("Should repeat wrapped wave the right number of times", func(t *testing.T) {
		wrapped := sound.NewWaveWithMaxDuration(&sound.MockWave{}, time.Millisecond)
		testwave := NewLoop(wrapped, time.Millisecond, 3)
		if _, err := testwave.Value(0, 3*time.Millisecond); errors.Is(err, sound.ErrEndOfWave) {
			t.Fatal("shouldn't get error end of wave yet")
		}
		if _, err := testwave.Value(0, 3*time.Millisecond+1); !errors.Is(err, sound.ErrEndOfWave) {
			t.Fatal("should get error end of wave")
		}
	})

	t.Run("Should restitute the wrapped wave's value without modification", func(t *testing.T) {
		wrapped := sound.NewWaveWithMaxDuration(&sound.MockWave{}, time.Millisecond)
		testwave := NewLoop(wrapped, time.Millisecond, 1)
		checkAt := time.Millisecond - 1
		got, _ := testwave.Value(0, checkAt)
		want, _ := wrapped.Value(0, checkAt)
		if got != want {
			t.Fatalf("want %f but got %f", want, got)
		}
	})
}
