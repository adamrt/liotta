package main

import (
	"fmt"
	"image"
)

const (
	// Image
	AspectRatio = 16.0 / 9.0
	ImageWidth  = 400
	ImageHeight = int(ImageWidth / AspectRatio)

	SamplesPerPixel = 100
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, ImageWidth, ImageHeight))

	world := World{
		Sphere{Vec3{0, 0, -1}, 0.5},
		Sphere{Vec3{0, -100.5, -1}, 100},
	}

	camera := NewCamera()

	for y := 0; y < ImageHeight; y++ {
		fmt.Printf("\rScanlines remaining: %d  ", ImageHeight-y)
		for x := 0; x < ImageWidth; x++ {

			pixelColor := Vec3{0, 0, 0}
			for s := 0; s < SamplesPerPixel; s++ {
				u := (float64(x) + random()) / float64(ImageWidth-1.0)
				v := (float64(y) + random()) / float64(ImageHeight-1.0)
				ray := camera.GetRay(u, v)
				pixelColor = pixelColor.Add(ray.Color(world))

			}

			rgba := pixelColor.RGBA(SamplesPerPixel)

			inverted_y := (ImageHeight - 1) - y
			img.Set(x, inverted_y, rgba)
		}
	}

	fmt.Printf("\nDone.\n")
	writePNG(img, "output.png")
}
