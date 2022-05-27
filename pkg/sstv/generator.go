package sstv

import (
	"fmt"
	"image"
	"time"

	"github.com/DylanMeeus/GoAudio/wave"
	"github.com/po1yb1ank/go-sstv/pkg/audio"
)

type Generator int

const (
	MartinM1 Generator = iota
	MartinM2
	Robot36
	Robot72
	ScottieS1
	ScottieS2
)

type IGenerator interface {
	generateFromImage(image.Image) error
}

type SSTVGenerator struct {
	generator  IGenerator
	duration   time.Duration
	sampleRate int
	minX       int
	minY       int
	vis        int
	audioUnit  *audio.AudioUnit
}

func (g *SSTVGenerator) GenerateWaveFromImage(i image.Image) error {
	if !g.isCorrectSize(i) {
		return fmt.Errorf("incorrect picture size. Should be: width >= %v, height >= %v", g.minX, g.minY)
	}
	g.generator.generateFromImage(i)

	return nil

}

func (g *SSTVGenerator) isCorrectSize(i image.Image) bool {
	return i.Bounds().Dx() >= g.minX && i.Bounds().Dy() >= g.minY
}

func NewSSTVGenerator(g Generator, sampleRate int) (*SSTVGenerator, error) {
	aunit, err := audio.NewSineAudioUnit(sampleRate)
	if err != nil {
		return nil, err
	}

	switch g {
	case MartinM1:
		return &SSTVGenerator{
			generator: &martinM1Gen{},
			duration:  0,
			minX:      320,
			minY:      256,
			audioUnit: aunit,
		}, nil
	case MartinM2:
		return &SSTVGenerator{
			generator: &martinM1Gen{},
			duration:  0,
			minX:      160,
			minY:      256,
			audioUnit: aunit,
		}, nil
	case Robot36:
		return &SSTVGenerator{
			generator: &martinM1Gen{},
			duration:  0,
			minX:      320,
			minY:      240,
			audioUnit: aunit,
		}, nil
	case Robot72:
		return &SSTVGenerator{
			generator: &martinM1Gen{},
			duration:  0,
			minX:      320,
			minY:      240,
			audioUnit: aunit,
		}, nil
	case ScottieS1:
		return &SSTVGenerator{
			generator: &martinM1Gen{},
			duration:  0,
			minX:      320,
			minY:      256,
			audioUnit: aunit,
		}, nil
	case ScottieS2:
		return &SSTVGenerator{
			generator: &martinM1Gen{},
			duration:  0,
			minX:      160,
			minY:      256,
			audioUnit: aunit,
		}, nil
	default:
		return &SSTVGenerator{}, fmt.Errorf("wrong generator")
	}
}

func (g *SSTVGenerator) generateVISSignal(vis int) []wave.Frame {
	frames := []wave.Frame{}
	// Leader tone
	frames = append(frames, g.audioUnit.GenerateFrame(0.3, 1900))
	// Break
	frames = append(frames, g.audioUnit.GenerateFrame(0.01, 0))
	// LeaderTone
	frames = append(frames, g.audioUnit.GenerateFrame(0.3, 1900))
	//VIS start bit
	frames = append(frames, g.audioUnit.GenerateFrame(0.03, 1200))

}
