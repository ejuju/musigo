# Quick notes

## Run tests

This section describes how to run tests for this project.

### Check code (unit tests and linting)
```
$ ./scripts/check_code.sh
```

### Show covered and uncovered code
```
$ go test -coverprofile cover.out ./...
$ go tool cover -html=cover.out
```

## Ideas for later

- [ ] Generate midi notes and chords variables with code generation
- [ ] Separate file encoder from player
- [ ] Add feature to decode PCM / MP3 / WAV into a playable wave
- [ ] Add script to build / play song file from docker container (so you don't have to install players)
- [ ] Improve perf, add concurrency to improve wave to pulse conversion calculations and decoding / encoding
- [ ] Implement filters by checking  the amplitude and adjusting the float64 in consequence
- [ ] Have instrument tracks send their waves by chunk to a central engine
- [ ] Web GUI / server