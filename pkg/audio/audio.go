package audio

import (
	"github.com/DylanMeeus/GoAudio/synthesizer"
	"github.com/DylanMeeus/GoAudio/wave"
)

type AudioUnit struct {
	sampleRate int
	oscillator *synthesizer.Oscillator
}

func NewSineAudioUnit(sr int) (*AudioUnit, error) {
	osc, err := synthesizer.NewOscillator(int(sr), synthesizer.SINE)
	if err != nil {
		return nil, err
	}
	return &AudioUnit{
		sampleRate: sr,
		oscillator: osc,
	}, nil
}

func (a *AudioUnit) GenerateFrame(duration, frequency float64) wave.Frame {
	return wave.Frame(
		synthesizer.ADSR(
			1,
			duration,
			0,
			0,
			duration,
			0,
			float64(a.sampleRate),
			int(duration*float64(a.sampleRate)),
		) * a.oscillator.Tick(frequency))
}
