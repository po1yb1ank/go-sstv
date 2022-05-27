package main

import (
	"log"

	"github.com/po1yb1ank/go-sstv/pkg/img"
	"github.com/po1yb1ank/go-sstv/pkg/sstv"
)

func main() {
	img, err := img.ScanImage("image.png")
	if err != nil {
		log.Fatal(err)
	}
	sstv.GenerateFromImage(img)
}
