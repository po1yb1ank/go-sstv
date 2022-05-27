package img

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"os"
)

type ImageIterator struct {
	img         image.Image
	currentLine int
	boundX      int
	boundY      int
}

func NewImageIterator(img image.Image) *ImageIterator {
	return &ImageIterator{
		img:         img,
		currentLine: -1,
		boundX:      img.Bounds().Max.X,
		boundY:      img.Bounds().Max.Y,
	}
}

func (i *ImageIterator) NextLineColors() (colors []color.Color, err error) {
	i.currentLine++

	if i.currentLine > i.boundY {
		return nil, fmt.Errorf("out of bounds")
	}

	for dx := 0; dx <= i.boundX; dx++ {
		colors = append(colors, i.img.At(dx, i.currentLine))
	}

	return colors, nil
}

func ScanImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	return img, err
}
