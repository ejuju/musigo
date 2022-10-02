package audio

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

func TestFFplayPlayer(t *testing.T) {
	t.Parallel()

	t.Run("Should implement the Player interface", func(t *testing.T) {
		var _ Player = &FFPlayPlayer{}
	})

	t.Run("Should validate required inputs", func(t *testing.T) {
		validWave := sound.NewWaveWithMaxDuration(&sound.MockWave{}, time.Second)
		var invalidWave sound.Wave = nil
		validSampleRate := 1
		invalidSampleRate := 0

		tests := []struct {
			wave       sound.Wave
			sampleRate int
			wantErr    bool
		}{
			{wave: validWave, sampleRate: validSampleRate, wantErr: false},  // check both valid
			{wave: invalidWave, sampleRate: validSampleRate, wantErr: true}, // check one invalid
			{wave: validWave, sampleRate: invalidSampleRate, wantErr: true}, // check one invalid
		}

		for i, test := range tests {
			player := &FFPlayPlayer{
				Wave:       test.wave,
				SampleRate: test.sampleRate,
				noExec:     true,
			}
			err := player.Play()
			if (err != nil) != test.wantErr {
				// make sure the error is not ErrFFPlayCommand as this error is expected in this case.
				if !errors.Is(err, ErrFFPlayCommand) {
					t.Fatalf("unexpected error at index %d, wantErr is %v but got %v", i, test.wantErr, err)
				}
			}
		}
	})

	t.Run("Should use the right command to launch ffplay", func(t *testing.T) {
		sampleRate := 34567
		filename := "foo_command"
		got := newFFPlayCommand(sampleRate, filename)

		// command start with ffplay
		if !strings.HasPrefix(got, "ffplay"+" ") {
			t.Fatalf("Want a string that starts with \"ffplay\" but got \"%s\"", got)
		}

		// command should end with filename
		if !strings.HasSuffix(got, " "+filename) {
			t.Fatalf("Want a string that ends with filename \"%s\" but got \"%s\"", filename, got)
		}

		// command should include the right frame format with -f flag
		if !strings.Contains(got, " -f f64le ") {
			t.Fatalf("Want a string that includes the frame format flag but got \"%s\"", got)
		}

		// command should include sample rate with -ar flag
		if !strings.Contains(got, " -ar "+strconv.Itoa(sampleRate)+" ") {
			t.Fatalf("Want a string that includes the sample rate flag but got: \"%s\"", got)
		}

		// command should include -autoexit flag
		if !strings.Contains(got, " -autoexit ") {
			t.Fatalf("Want a string that includes the autoexit flag but got: \"%s\"", got)
		}

		// command should use showmode 1
		if !strings.Contains(got, " -showmode 1 ") {
			t.Fatalf("Want a string that includes the showmode 1 flag but got: \"%s\"", got)
		}
	})

	t.Run("Should save file only if desired", func(t *testing.T) {
		filepath := "foo_save_file"
		player := &FFPlayPlayer{
			Wave:       sound.NewWaveWithMaxDuration(&sound.MockWave{}, time.Second),
			SampleRate: 44100,
			Filepath:   filepath,
			noExec:     true,
		}

		// play without saving file
		err := player.Play()
		if err != nil && !errors.Is(err, ErrFFPlayCommand) {
			// make sure the error is not ErrFFPlayCommand as this error is expected in this case.
			t.Fatal("unexpected error", err)
		}

		// make sure no file is found
		_, err = os.Open(filepath)
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("want error %s but got %s", os.ErrExist, err)
		}

		// play with saving file
		player.SaveFile = true
		err = player.Play()
		if err != nil && !errors.Is(err, ErrFFPlayCommand) {
			// make sure the error is not ErrFFPlayCommand as this error is expected in this case.
			t.Fatal("unexpected error:", err)
		}

		// make sure the file is found
		_, err = os.Open(filepath)
		if errors.Is(err, os.ErrNotExist) {
			t.Fatalf("expected no error but got err %s", err)
		}

		// cleanup: remove saved file if it was created
		err = os.Remove(filepath)
		if err != nil {
			t.Logf("warning: failed to remove test file: %s (err: %s)", filepath, err)
		}
	})
}
