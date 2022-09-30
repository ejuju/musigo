package audio

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
)

type PCMEncoder struct{}

// EncodePCM encodes a sound's PCM representation to an io.Writer
func (e *PCMEncoder) Encode(w io.Writer, frames []float64) error {
	if len(frames) == 0 {
		return errors.New("no frames were provided")
	}
	if w == nil {
		return errors.New("no io.Writer was provided")
	}

	// encode each pulse to writer
	for _, pulse := range frames {
		var buf [8]byte
		binary.LittleEndian.PutUint64(buf[:], math.Float64bits(pulse))
		_, err := w.Write(buf[:])
		if err != nil {
			return err
		}
	}

	return nil
}

type PCMDecoder struct{}

func (e *PCMDecoder) Decode(r io.Reader) (*Sample, error) {
	if r == nil {
		return nil, errors.New("no io.Reader was provided")
	}
	return nil, nil
}
