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

### Folder structure

- `pkg/audio`: Audio encoding and decoding (PCM, WAV, MIDI, etc.) and players
- `pkg/maths`: Math utilities (floats utilities, noise/random functions, interpollation functions, etc.)
- `pkg/music`: Musical primitives (notes, chords, scales, tempo, composition, etc.)
- `pkg/sound`: Sound synthesis (oscillators, waves, envelopes, effects, etc.)
- `pkg/musigo`: Musigo debugging layer.

### Features for v1

- Sound:
    - [x] Basic oscillators (sine, square, sawtooth)
    - [ ] Amplitude envelope
    - [ ] Patterns (to build progressions)
- Audio:
    - [x] PCM encoding
    - [x] PCM playing (with ffplay)
    - [ ] PCM decoding
    - [ ] WAV encoding
    - [ ] WAV decoding
