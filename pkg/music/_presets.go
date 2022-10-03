package music

import (
	"time"

	"github.com/ejuju/musigo/pkg/sound"
)

func NewPresetMetronome(freq float64) sound.Wave {
	amplCtrlWave := sound.NewControlWave(nil, false, []*sound.ControlWaveSegment{
		{Duration: 10 * time.Millisecond, StartValue: 1.0, EndValue: 1.0},
		{Duration: 10 * time.Millisecond},
	})

	waves := []sound.Wave{}
	multiples := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	for _, multiple := range multiples {
		multwave := sound.NewSynthWave(&sound.Sine{}, freq*multiple)
		waves = append(waves, sound.NewAmplitudeEnvelope(multwave, amplCtrlWave))
	}

	return sound.NewMergedWaves(waves...)
}
