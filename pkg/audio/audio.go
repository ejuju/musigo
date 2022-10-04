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
func Frames(wave sound.Wave, sampleRate int, startOffset, duration time.Duration) ([]float64, error) {
	if wave == nil {
		return nil, errors.New("wave is not defined, unable to get frames")
	}
	if sampleRate <= 0 {
		return nil, fmt.Errorf("invalid sample rate: %d, sample rate should be positive", sampleRate)
	}

	frames := []float64{}
	step := float64(time.Second) / float64(sampleRate)
	for i := float64(startOffset); i < float64(startOffset+duration); i += step {
		val, err := wave.Value(time.Duration(i))
		if err != nil {
			return nil, err
		}

		frames = append(frames, val)
	}

	return frames, nil
}
