package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const Width = 256
const Height = 256

func main() {
	img := image.NewRGBA(image.Rect(0, 0, Height, Width))

	for j := 0; j < Height; j++ {
		fmt.Printf("\rScanlines remaining: %d  ", Height-j)
		for i := 0; i < Width; i++ {
			r := float64(i) / (Width - 1)
			g := float64(j) / (Height - 1)
			b := 0.25

			// j calculation is to invert the image.
			img.Set(i, (Height-1)-j, color.RGBA{
				uint8(255.999 * r),
				uint8(255.999 * g),
				uint8(255.999 * b),
				255,
			})
		}
	}

	fmt.Printf("\nDone.\n")

	f, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
