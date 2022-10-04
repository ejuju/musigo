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

func main() {
	var bpm = music.BPM(120)

	chords := [][]float64{
		music.ChordMajor7.Octave(0, 1, 2).Frequencies(music.NoteA3.Hertz())[:8],
		music.ChordMinor.Octave(0, 1).Frequencies(music.NoteA4.Hertz())[:8],
		music.ChordMinor7.Octave(0, 1).Frequencies(music.NoteA4.Hertz())[:8],
		music.ChordMajor.Octave(0, 1).Frequencies(music.NoteC4.Hertz())[:8],
	}

	tracks := musigo.Tracks{
		"arpeggio": musigo.NewTrack(&sound.Sine{}, func(c *musigo.Controller) {
			for _, notes := range chords {
				for _, note := range notes {
					c.Play(bpm.Beats(0.5), []sound.Effect{
						sound.NewAmplitudeEnvelope(sound.NewControlWave(nil, []*sound.ControlWaveSegment{
							{Duration: bpm.Beats(0.25), EndValue: 1.0},
							{Duration: bpm.Beats(0.25)},
						})),
					}, note)
				}
			}
		}),
		"bassline": musigo.NewTrack(&sound.SawTooth{}, func(c *musigo.Controller) {
			for i := 0; i < 16; i++ {
				for _, notes := range chords {
					c.Play(bpm.Beats(1), []sound.Effect{
						sound.NewAmplitudeEnvelope(sound.NewControlWave(nil, []*sound.ControlWaveSegment{
							{Duration: bpm.Beats(0.5), EndValue: 1.0},
							{Duration: bpm.Beats(0.5)},
						})),
					}, notes[0])
				}
			}
		}),
	}

	err := audio.FFPlayPlayer{Wave: tracks.Merge(), SampleRate: 44100}.Play()
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
- Music:
    - [x] Handle loops
    - [ ] Provide notes
    - [ ] Provide chords and scales
    - [ ] Handle arpeggios
    - [ ] Handle tempo