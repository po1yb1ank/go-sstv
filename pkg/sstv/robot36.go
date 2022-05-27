package sstv

import (
	"image"

	"github.com/po1yb1ank/go-sstv/pkg/img"
)

type robot36Gen struct {
	vis int
}

func (g *robot36Gen) generateFromImage(i image.Image, sr int) error {
	imgIterator := img.NewImageIterator(i)

	for line, err := imgIterator.NextLineColors(); err == nil; {
		for _, color := range line {

		}
	}

}
