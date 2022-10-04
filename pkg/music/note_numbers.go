package music

// NoteNumber represents the MIDI notation compliant number attributed to notes.
type NoteNumber int

// Hz returns the frequency (in hertz) corresponding to the given note.
func (n NoteNumber) Hz() float64 {
	return RelativeFrequency(440.00, float64(n)-69.00)
}

// Common notes defined as ready to use variables for convenience.

const (
	NoteA0 NoteNumber = 21 + iota
	NoteBb0
	NoteB0

	NoteC1
	NoteDb1
	NoteD1
	NoteEb1
	NoteE1
	NoteF1
	NoteGb1
	NoteG1
	NoteAb1
	NoteA1
	NoteBb1
	NoteB1

	NoteC2
	NoteDb2
	NoteD2
	NoteEb2
	NoteE2
	NoteF2
	NoteGb2
	NoteG2
	NoteAb2
	NoteA2
	NoteBb2
	NoteB2

	NoteC3
	NoteDb3
	NoteD3
	NoteEb3
	NoteE3
	NoteF3
	NoteGb3
	NoteG3
	NoteAb3
	NoteA3
	NoteBb3
	NoteB3

	NoteC4
	NoteDb4
	NoteD4
	NoteEb4
	NoteE4
	NoteF4
	NoteGb4
	NoteG4
	NoteAb4
	NoteA4
	NoteBb4
	NoteB4

	NoteC5
	NoteDb5
	NoteD5
	NoteEb5
	NoteE5
	NoteF5
	NoteGb5
	NoteG5
	NoteAb5
	NoteA5
	NoteBb5
	NoteB5

	NoteC6
	NoteDb6
	NoteD6
	NoteEb6
	NoteE6
	NoteF6
	NoteGb6
	NoteG6
	NoteAb6
	NoteA6
	NoteBb6
	NoteB6

	NoteC7
	NoteDb7
	NoteD7
	NoteEb7
	NoteE7
	NoteF7
	NoteGb7
	NoteG7
	NoteAb7
	NoteA7
	NoteBb7
	NoteB7

	NoteC8
	NoteDb8
	NoteD8
	NoteEb8
	NoteE8
	NoteF8
	NoteGb8
	NoteG8
	NoteAb8
	NoteA8
	NoteBb8
	NoteB8

	NoteC9
	NoteDb9
	NoteD9
	NoteEb9
	NoteE9
	NoteF9
	NoteGb9
	NoteG9
	NoteAb9
)
