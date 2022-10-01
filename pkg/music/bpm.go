package music

import "time"

// BPM corresponds to the number of beats per minute.
// It is used to represent a tempo.
type BPM float64

func (bpm BPM) BeatsToDuration(beats float64) time.Duration {
	return time.Duration(float64(time.Minute) * beats / float64(bpm))
}
