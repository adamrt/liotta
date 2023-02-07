package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/cheggaaa/pb/v3"
)

const (
	// Image
	aspectRatio = 4.0 / 3.0
	ImageWidth  = 800
	ImageHeight = int(ImageWidth / aspectRatio)

	SamplesPerPixel = 100
	MaxDepth        = 50
	Threads         = 12
)

type Pixel struct {
	x, y  int
	color color.RGBA
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, ImageWidth, ImageHeight))
	world := randomScene()

	vfov := 20.0
	camPos := Vec3{13, 2, 3}
	lookAt := Vec3{0, 0, 0}
	up := Vec3{0, 1, 0}
	distToFocus := 10.0 // camPos.Sub(lookAt).Length()
	aperture := 0.1     // 2.0

	camera := NewCamera(camPos, lookAt, up, vfov, aspectRatio, aperture, distToFocus)

	bar := pb.StartNew(ImageHeight * ImageWidth)
	pixelChan := make(chan Pixel)

	// Start a new renderer for each thread
	for i := 0; i < Threads; i++ {
		go func() {
			for pixel := range pixelChan {
				var color Vec3
				for s := 0; s < SamplesPerPixel; s++ {
					u := (float64(pixel.x) + random()) / float64(ImageWidth)
					v := (float64(pixel.y) + random()) / float64(ImageHeight)
					ray := camera.GetRay(u, v)
					color = color.Add(ray.Color(world, MaxDepth))
				}
				img.Set(pixel.x, ImageHeight-1-pixel.y, color.RGBA(SamplesPerPixel))
				bar.Increment()
			}
		}()
	}

	// Populate the pixels channel with each pixel
	for y := 0; y < ImageHeight; y++ {
		for x := 0; x < ImageWidth; x++ {
			pixelChan <- Pixel{x: x, y: y}
		}
	}
	close(pixelChan)

	fmt.Printf("\nDone.\n")
	writePNG(img, "output.png")
	bar.Finish()
}

func randomScene() World {
	world := World{}

	// Ground sphere
	materialGround := NewLambert(Vec3{0.5, 0.5, 0.5})
	world = append(world, NewSphere(Vec3{0, -1000, 0}, 1000, materialGround))

	// Random spheres
	xx := 11
	for a := -xx; a < xx; a++ {
		for b := -xx; b < xx; b++ {
			choose_mat := random()
			center := Vec3{float64(a) + 0.9*random(), 0.2, float64(b) + 0.9*random()}

			if center.Sub(Vec3{4, 0.2, 0}).Length() > 0.9 {
				if choose_mat < 0.8 {
					// diffuse
					albedo := Vec3Rand(0, 1).Mul(Vec3Rand(0, 1))
					material := NewLambert(albedo)
					world = append(world, NewSphere(center, 0.2, material))
				} else if choose_mat < 0.95 {
					// metal
					albedo := Vec3{randomF(0.5, 1), randomF(0.5, 1), randomF(0.5, 1)}
					fuzz := randomF(0, 0.5)
					material := NewMetal(albedo, fuzz)
					world = append(world, NewSphere(center, 0.2, material))
				} else {
					// glass
					material := NewDielectric(1.5)
					world = append(world, NewSphere(center, 0.2, material))
				}
			}
		}
	}

	// Three large constant spheres

	material1 := NewDielectric(1.5)
	world = append(world, NewSphere(Vec3{0, 1, 0}, 1.0, material1))

	material2 := NewLambert(Vec3{0.4, 0.2, 0.1})
	world = append(world, NewSphere(Vec3{-4, 1, 0}, 1.0, material2))

	material3 := NewMetal(Vec3{0.7, 0.6, 0.5}, 0.0)
	world = append(world, NewSphere(Vec3{4, 1, 0}, 1.0, material3))

	return world
}
