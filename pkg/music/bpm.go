package music

import "time"

// BPM corresponds to the number of beats per minute.
// It is used to represent a tempo.
type BPM float64

// Duration returns the time.Duration corresponding to a certain number of beats.
func (bpm BPM) Time(numBeats float64) time.Duration {
	return time.Duration(float64(time.Minute) * numBeats / float64(bpm))
}
