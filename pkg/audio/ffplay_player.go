package audio

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

// FFplayPlayer uses ffplay to play the provided frames.
// It produces a .pcm file under the hood to encode the output sound wave.
// This file can be saved by setting the SaveFile field to true.
type FFPlayPlayer struct {
	Wave                sound.Wave
	SampleRate          int
	DurationStartOffset time.Duration
	Duration            time.Duration
	Filepath            string
	SaveFile            bool
	noExec              bool // to prevent ffplay command execution
}

var ErrFFPlayCommand = errors.New("failed to execute ffplay command")

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
	if p.Duration <= 0 {
		return fmt.Errorf("invalid duration: %s", p.Duration)
	}

	// get output frames
	frames, err := Frames(p.Wave, p.SampleRate, p.DurationStartOffset, p.Duration)
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

	// the call to remove the file is deferred (in case the command execution fails, then the the file wouldn't be removed)
	defer func() error {
		// remove file after play if desired
		if !p.SaveFile {
			err := os.Remove(p.Filepath)
			if err != nil {
				return fmt.Errorf("failed to delete file (%s): %w", p.Filepath, err)
			}
		}
		return nil
	}()

	// allow caller to prevent command execution for testing for example.
	if p.noExec {
		return nil
	}

	// Read output file with ffplay (by launching ffplay from the CLI)
	cmdstr := strings.Split(newFFPlayCommand(p.SampleRate, p.Filepath), " ")
	err = exec.Command(cmdstr[0], cmdstr[1:]...).Run()
	if err != nil {
		return fmt.Errorf("failed to play PCM file using ffplay: %w: %s", ErrFFPlayCommand, err.Error())
	}
	return nil
}

// newFFPlayCommand returns the command string used to play a PCM file with ffplay.
func newFFPlayCommand(sampleRate int, filepath string) string {
	return "ffplay" + " " +
		"-f f64le" + " " +
		"-ar " + strconv.Itoa(sampleRate) + " " +
		"-autoexit" + " " +
		"-showmode 1" + " " +
		filepath
}
