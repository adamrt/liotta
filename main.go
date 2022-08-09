package main

import (
	"fmt"
	"image"
)

const (
	// Image
	aspectRatio = 4.0 / 3.0
	ImageWidth  = 1200
	ImageHeight = int(ImageWidth / aspectRatio)

	SamplesPerPixel = 500
	MaxDepth        = 50

	// There are 3 diffuse methods to, this is to switch between them.
	LambertDiffuseMethod = 2
)

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
