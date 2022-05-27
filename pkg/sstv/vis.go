package sstv

import (
	"github.com/DylanMeeus/GoAudio/synthesizer"
	"github.com/DylanMeeus/GoAudio/wave"
)

func generateVISSignal(vis int) []wave.Frame {
	frames := []wave.Frame{}
	// Leader tone
	frames = append(frames, wave.Frame(synthesizer.ADSR(1, 0.3, 0, 0, 0.3, 0, float64(sr), int(0.3*float64(sr)))*osc.Tick(1900)))
	// Break
	frames = append(frames, audio.generateFrame)

}
