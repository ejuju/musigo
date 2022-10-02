package audio

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

// Encoder encodes audio frames to a writer.
type Encoder interface {
	Encode(w io.Writer, frames []float64) error
}

// Decoder decodes an audio file to a sound.Wave that can be used with your code.
type Decoder interface {
	Decode(r io.Reader) (*Sample, error)
}

// Player plays audio.
type Player interface {
	Play() error
}

// Frames returns audio frames generated using the provided sound wave.
// The provided sample rate is in number of frames per second (= hertz).
func Frames(wave sound.Wave, sampleRate int) ([]float64, error) {
	if wave == nil {
		return nil, errors.New("wave is not defined, unable to get frames")
	}
	if sampleRate <= 0 {
		return nil, fmt.Errorf("invalid sample rate: %d, sample rate should be positive", sampleRate)
	}

	frames := []float64{}
	for i := 0; true; i++ {
		x := float64(int(time.Second)*i) / float64(sampleRate) // current elapsed time passed to wave

		val, err := wave.Value(time.Duration(x))
		if err != nil {
			if errors.Is(err, sound.ErrEndOfWave) {
				break
			}
			return nil, err
		}

		frames = append(frames, val)
	}

	return frames, nil
}
