# Make music with Go

### What is Musigo?

Musigo is an audio synthesis toolkit for Go developers. 
You can use it to compose algorithmic music.

### Requirements

1. Get the Golang Musigo module in your project `go get -u github.com/ejuju/musigo`
1. Make sure `ffplay` is installed and can be called from the command line.

### Getting started

Copy the code in `examples/0_hello_world/main.go` and run it with `go run .`.
Once you can hear the output sound, you're good to go, you can just jump through the examples one by one.

### Design

Sound is made by combining together waves.
You can generate oscillations (sine, square, etc.) that you combine with other waves to build more complex sounds.
Once these sounds have been built, you can build patterns with them to map them through time.
Sounds can be compiled to audio frames (for PCM encoding) and combined.
They are saved to local files once they are compiled and can be reloaded later.
The idea is to build the smallest file possible for each possible sound you need in your song. This could be a bassline, a kick, a snare, a ambient noise, a chord progression, guitar solo, etc.

### Folder structure

- `pkg/audio`: Audio encoding and decoding (PCM, WAV, MIDI, etc.) and players
- `pkg/maths`: Math utilities (floating point comparison utilities, noise/random functions, interpollation functions, etc.)
- `pkg/music`: Musical primitives (notes, chords, scales, tempo, composition, etc.)
- `pkg/sound`: Sound synthesis (oscillators, waves, envelopes, effects, etc.)
- `pkg/musigo`: Musigo debugging layer.


