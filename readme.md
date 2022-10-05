# Make music with Go

### What is Musigo?

Musigo is an audio synthesis toolkit for Go developers. 
You can use it to compose algorithmic music.

### Requirements

1. Add Musigo to your Go project with `go get -u github.com/ejuju/musigo`
1. Make sure `ffplay` is installed and can be called from the command line.

### Getting started

Copy the code in `examples/0_hello_world/main.go` and run it with `go run .`.
Once you can hear the output sound, you're good to go, you can just jump through the examples one by one.

Basic example:
```go
// main.go
package main

import (
	"github.com/ejuju/musigo/pkg/audio"
	"github.com/ejuju/musigo/pkg/music"
	"github.com/ejuju/musigo/pkg/musigo"
	"github.com/ejuju/musigo/pkg/sound"
)

var (
	bpm    = music.BPM(120)
	key    = music.NoteA4
	chords = []music.Notes{
		music.ChordMajor7.Notes(key.Hz()).Repeat(2)[:8],
		music.ChordMinor.Notes((key + 2).Hz()).Repeat(2)[:8],
		music.ChordMinor.Notes((key + 2).Hz()).Repeat(2)[:8],
		music.ChordMinor7.Notes((key + 2).Hz()).Repeat(2)[:8],
	}
)

var basslineTrack = musigo.NewTrack(sound.SawTooth{}, func(c *musigo.Controller) {
		c.Play(bpm.Beats(6), nil, chords[0].Octave(-3)[0])
		c.Wait(bpm.Beats(2))
},
	sound.NewAmplitudeEnvelope(sound.NewControlWave(nil, 0).
		Append(bpm.Beats(16), 0.1).
		Append(bpm.Beats(32 * 4), 0.1).
		Append(bpm.Beats(8), 0).
		Append(bpm.Beats(8), 0),
)

var chordsTrack = musigo.NewTrack(sound.Sine{}, func(c *musigo.Controller) {
		for _, notes := range chords {
			c.Play(bpm.Beats(8), nil, notes...)
		}
},
	sound.NewAmplitudeEnvelope(sound.NewControlWave(nil, 0).
		Append(bpm.Beats(1), 0.6).
		Append(bpm.Beats(7), 0),
)

func main() {
	tracks := musigo.Tracks{
		"bassline": basslineTrack,
		"chords":   chordsTrack,
	}

	player := audio.FFPlayPlayer{
		Wave:       tracks.Merge(),
		Duration:   bpm.Beats(32),
		SampleRate: 44100,
	}

	err := player.Play()
	if err != nil {
		panic(err)
	}
}
```

### Folder structure

- `pkg/audio`: Audio encoding and decoding (PCM, WAV, MIDI, etc.) and players
- `pkg/maths`: Math utilities (floats utilities, noise/random functions, interpollation functions, etc.)
- `pkg/music`: Musical primitives (notes, chords, scales, tempo, composition, etc.)
- `pkg/sound`: Sound synthesis (oscillators, waves, envelopes, effects, etc.)

### Features for v1

- Sound:
    - [x] Provide basic oscillators (sine, square, sawtooth)
    - [x] Control amplitude over time (ex: ADSR envelope)
    - [x] Control frequency over time (ex: melody)
- Audio:
    - [x] Encode `sound.Wave` as PCM
    - [x] Play PCM files (with ffplay)
    - [ ] Decode PCM to `sound.Wave`
    - [ ] Decode WAV to PCM
    - [ ] Encode PCM to WAV
    - [ ] Provide percussion audio samples (with go embed)
	- [ ] Add live player for real-time audio processing
- Music:
    - [x] Handle loops
    - [ ] Provide notes
    - [ ] Provide chords and scales
    - [ ] Handle arpeggios
    - [ ] Handle tempo