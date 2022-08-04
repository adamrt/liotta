package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

const (
	// Image
	AspectRatio = 16.0 / 9.0
	ImageWidth  = 400
	ImageHeight = int(ImageWidth / AspectRatio)

	// Camera
	ViewportHeight = 2.0
	ViewportWidth  = AspectRatio * ViewportHeight
	FocalLength    = 1.0
)

func main() {

	var (
		origin          = Vec3{0, 0, 0}
		horizontal      = Vec3{ViewportWidth, 0, 0}
		vertical        = Vec3{0, ViewportHeight, 0}
		lowerLeftCorner = origin.Sub(horizontal.ScalarDiv(2)).Sub(vertical.ScalarDiv(2)).Sub(Vec3{0, 0, FocalLength})
	)

	img := image.NewRGBA(image.Rect(0, 0, ImageWidth, ImageHeight))

	for j := 0; j < ImageHeight; j++ {
		fmt.Printf("\rScanlines remaining: %d  ", ImageHeight-j)
		for i := 0; i < ImageWidth; i++ {

			u := float64(i) / float64(ImageWidth-1)
			v := float64(j) / float64(ImageHeight-1)

			direction := lowerLeftCorner.Add(horizontal.ScalarMul(u)).Add(vertical.ScalarMul(v)).Sub(origin)
			ray := Ray{
				origin:    origin,
				direction: direction,
			}

			// j calculation is to invert the image.
			img.Set(i, (ImageHeight-1)-j, ray.Color())
		}
	}

	fmt.Printf("\nDone.\n")
	writePNG(img, "output.png")
}

func writePNG(img *image.RGBA, filename string) {
	f, err := os.Create(filename)
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
