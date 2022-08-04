package main

import (
	"fmt"
	"image"
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
			pixel := Vec3{
				x: float64(i) / (Width - 1),
				y: float64(j) / (Height - 1),
				z: 0.25,
			}

			// j calculation is to invert the image.
			img.Set(i, (Height-1)-j, pixel.RGBA())

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
