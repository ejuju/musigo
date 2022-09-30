package audio

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

// FFplayPlayer uses ffplay to play the provided frames.
// It produces a .pcm file under the hood to encode the output sound wave.
// This file can be saved by setting the SaveFile field to true.
type FFPlayPlayer struct {
	Wave       sound.Wave
	SampleRate int
	Filepath   string
	SaveFile   bool
	Freq       float64
}

func (p FFPlayPlayer) Play() error {
	if p.SampleRate <= 0 {
		return fmt.Errorf("invalid sample rate: %d, sample rate must be positive", p.SampleRate)
	}
	if p.Filepath == "" {
		p.Filepath = strconv.Itoa(int(time.Now().Unix())) + ".pcm"
	}
	if p.Wave == nil {
		return errors.New("no wave was provided")
	}

	// get output frames
	frames, err := Frames(p.Freq, p.Wave, p.SampleRate)
	if err != nil {
		return fmt.Errorf("failed to get output frames from wave: %w", err)
	}

	// Create output file
	f, err := os.Create(p.Filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	// Encode PCM output to file
	err = (&PCMEncoder{}).Encode(f, frames)
	if err != nil {
		return fmt.Errorf("failed to encode PCM pulses: %w", err)
	}

	// Read output file with ffplay
	err = exec.CommandContext(context.Background(),
		"ffplay",
		"-f", "f64le",
		"-ar", strconv.Itoa(p.SampleRate),
		"-window_title", "Musigo - "+p.Filepath,
		"-autoexit",
		"-showmode", "1",
		p.Filepath,
	).Run()

	if err != nil {
		return fmt.Errorf("failed to play PCM file using ffplay: %w", err)
	}

	// remove file after play if desired
	if !p.SaveFile {
		err := os.Remove(p.Filepath)
		if err != nil {
			return fmt.Errorf("failed to delete file (%s): %w", p.Filepath, err)
		}
	}

	return nil
}
