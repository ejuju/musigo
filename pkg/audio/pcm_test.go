package audio

import (
	"bytes"
	"io"
	"testing"
)

func TestPCMEncoder(t *testing.T) {
	t.Parallel()

	t.Run("Should implement the Encoder interface", func(t *testing.T) {
		var _ Encoder = &PCMEncoder{}
	})

	t.Run("Should validate inputs", func(t *testing.T) {
		encoder := &PCMEncoder{}
		validFrames := []float64{1.0}
		invalidFrames := []float64{}
		validWriter := bytes.NewBuffer([]byte{})
		var invalidWriter io.Writer = nil

		tests := []struct {
			writer  io.Writer
			frames  []float64
			wantErr bool
		}{
			{writer: validWriter, frames: validFrames, wantErr: false},
			{writer: invalidWriter, frames: invalidFrames, wantErr: true},
			{writer: validWriter, frames: invalidFrames, wantErr: true},
			{writer: invalidWriter, frames: validFrames, wantErr: true},
		}

		for i, test := range tests {
			err := encoder.Encode(test.writer, test.frames)
			if (err != nil) != test.wantErr {
				t.Fatalf("unexpected error at index %d, wantErr is %v but got %v", i, test.wantErr, err)
			}
		}
	})
}

func TestPCMDecoder(t *testing.T) {
	t.Parallel()

	t.Run("Should implement the Decoder interface", func(t *testing.T) {
		var _ Decoder = &PCMDecoder{}
	})

	t.Run("Should validate required inputs", func(t *testing.T) {
		decoder := &PCMDecoder{}
		validReader := bytes.NewBuffer([]byte{})
		var invalidReader io.Reader = nil

		tests := []struct {
			reader  io.Reader
			wantErr bool
		}{
			{reader: validReader, wantErr: false},
			{reader: invalidReader, wantErr: true},
		}

		for i, test := range tests {
			_, err := decoder.Decode(test.reader)
			if (err != nil) != test.wantErr {
				// ignore command error
				t.Fatalf("unexpected error at index %d, wantErr is %v but got %v", i, test.wantErr, err)
			}
		}
	})
}
