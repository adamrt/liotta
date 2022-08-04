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
	MaxDepth        = 50

	// There are 3 diffuse methods to, this is to switch between them.
	LambertDiffuseMethod = 2
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, ImageWidth, ImageHeight))

	material_ground := Lambert{Vec3{0.8, 0.8, 0.0}}
	material_center := Lambert{Vec3{0.7, 0.3, 0.3}}
	material_left := Metal{Vec3{0.8, 0.8, 0.8}}
	material_right := Metal{Vec3{0.8, 0.6, 0.2}}

	world := World{
		NewSphere(Vec3{0.0, -100.5, -1.0}, 100.0, material_ground),
		NewSphere(Vec3{0.0, 0.0, -1.0}, 0.5, material_center),
		NewSphere(Vec3{-1.0, 0.0, -1.0}, 0.5, material_left),
		NewSphere(Vec3{1.0, 0.0, -1.0}, 0.5, material_right),
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
