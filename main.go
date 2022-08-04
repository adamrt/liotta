package main

import (
	"fmt"
	"image"
)

const (
	// Image
	aspectRatio = 16.0 / 9.0
	ImageWidth  = 400
	ImageHeight = int(ImageWidth / aspectRatio)

	SamplesPerPixel = 100
	MaxDepth        = 50

	// There are 3 diffuse methods to, this is to switch between them.
	LambertDiffuseMethod = 2
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, ImageWidth, ImageHeight))

	materialGround := NewLambert(Vec3{0.8, 0.8, 0.0})
	materialCenter := NewLambert(Vec3{0.1, 0.2, 0.5})
	materialLeft := NewDielectric(1.5)
	materialRight := NewMetal(Vec3{0.8, 0.6, 0.2}, 0.0)

	world := World{
		NewSphere(Vec3{0.0, -100.5, -1.0}, 100.0, materialGround),
		NewSphere(Vec3{0.0, 0.0, -1.0}, 0.5, materialCenter),
		NewSphere(Vec3{-1.0, 0.0, -1.0}, 0.5, materialLeft),
		NewSphere(Vec3{-1.0, 0.0, -1.0}, -0.45, materialLeft),
		NewSphere(Vec3{1.0, 0.0, -1.0}, 0.5, materialRight),
	}

	vfov := 90.0
	camPos := Vec3{-2, 1.0, 0.5}
	lookAt := Vec3{0, 0, -1}
	up := Vec3{0, 1, 0}
	camera := NewCamera(camPos, lookAt, up, vfov, aspectRatio)

	for y := 0; y < ImageHeight; y++ {
		fmt.Printf("\rScanlines remaining: %d  ", ImageHeight-y)
		for x := 0; x < ImageWidth; x++ {

			pixelColor := Vec3{0, 0, 0}
			for s := 0; s < SamplesPerPixel; s++ {
				u := (float64(x) + random()) / float64(ImageWidth-1.0)
				v := (float64(y) + random()) / float64(ImageHeight-1.0)
				ray := camera.GetRay(u, v)
				pixelColor = pixelColor.Add(ray.Color(world, MaxDepth))

			}

			rgba := pixelColor.RGBA(SamplesPerPixel)

			inverted_y := (ImageHeight - 1) - y
			img.Set(x, inverted_y, rgba)
		}
	}

	fmt.Printf("\nDone.\n")
	writePNG(img, "output.png")
}
