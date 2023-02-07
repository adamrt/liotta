package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randomF(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func random() float64 {
	return randomF(0.0, 1.0)
}

func clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func writePNG(img *image.RGBA, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}
