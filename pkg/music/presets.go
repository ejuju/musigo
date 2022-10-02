package music

// // multiples arg is optional.
// func NewPresetMetronome(freq float64) sound.Wave {
// 	amplCtrlWave := sound.NewControlWave(nil, []*sound.ControlWaveSegment{
// 		{Duration: 10 * time.Millisecond, StartValue: 1.0, EndValue: 1.0},
// 		{Duration: 10 * time.Millisecond},
// 	})
// 	rootWave := sound.NewWaveWithAmplitudeEnvelope(&sound.SineWave{}, amplCtrlWave)

// 	waves := []sound.Wave{rootWave}

// 	// add a bit of random noise
// 	waves = append(waves, sound.NewWaveWithAmplitudeEnvelope(
// 		sound.NewRandomWideBandNoise(0),
// 		sound.NewControlWave(nil, []*sound.ControlWaveSegment{
// 			{Duration: 5 * time.Millisecond, EndValue: 0.1},
// 			{Duration: 15 * time.Millisecond},
// 		})),
// 	)

// 	// add harmonics and overtones
// 	multiples := []float64{2, 3, 4, 5}
// 	for _, multiple := range multiples {
// 		multwave := sound.NewWaveWithFrequencyMultiplier(&sound.SineWave{}, multiple)
// 		waves = append(waves, sound.NewWaveWithAmplitudeEnvelope(multwave, amplCtrlWave))
// 	}

// 	return sound.NewMergedWaves(waves...)
// }
