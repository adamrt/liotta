package main

import (
	"image/color"
	"math"
)

func Vec3Rand(min, max float64) Vec3 {
	return Vec3{randomF(min, max), randomF(min, max), randomF(min, max)}
}

// Diffuse method 1
func Vec3RandomInUnitSphere() Vec3 {
	for {
		p := Vec3Rand(-1.0, 1.0)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

// Diffuse method 2
func Vec3RandUnit() Vec3 {
	return Vec3RandomInUnitSphere().Unit()
}

// Diffuse method 3
func Vec3RandomInHemisphere(normal Vec3) Vec3 {
	inUnitSphere := Vec3RandomInUnitSphere()
	// In the same hemisphere as the normal
	if inUnitSphere.Dot(normal) > 0.0 {
		return inUnitSphere
	} else {
		return inUnitSphere.Neg()
	}
}

type Vec3 struct {
	x, y, z float64
}

func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{
		x: v.x + u.x,
		y: v.y + u.y,
		z: v.z + u.z,
	}
}

func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{
		x: v.x - u.x,
		y: v.y - u.y,
		z: v.z - u.z,
	}
}

func (v Vec3) Mul(u Vec3) Vec3 {
	return Vec3{
		x: v.x * u.x,
		y: v.y * u.y,
		z: v.z * u.z,
	}
}

func (v Vec3) Div(u Vec3) Vec3 {
	return Vec3{
		x: v.x / u.x,
		y: v.y / u.y,
		z: v.z / u.z,
	}
}

func (v Vec3) ScalarMul(s float64) Vec3 {
	return Vec3{
		x: v.x * s,
		y: v.y * s,
		z: v.z * s,
	}
}

func (v Vec3) Neg() Vec3 {
	return Vec3{
		x: -v.x,
		y: -v.y,
		z: -v.z,
	}
}

func (v Vec3) ScalarDiv(s float64) Vec3 {
	return v.ScalarMul(1.0 / s)
}

func (v Vec3) Dot(u Vec3) float64 {
	return v.x*u.x + v.y*u.y + v.z*u.z
}

func (v Vec3) Cross(u Vec3) Vec3 {
	return Vec3{
		x: v.y*u.z - v.z*u.y,
		y: v.z*u.x - v.x*u.z,
		z: v.x*u.y - v.y*u.x,
	}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) Unit() Vec3 {
	return v.ScalarDiv(v.Length())
}

func (v Vec3) RGBA(samplesPerPixel int) color.RGBA {

	r := v.x
	g := v.y
	b := v.z

	// Divide the color by the number of samples and gamma-correct for
	// gamma=2.0.
	scale := 1.0 / float64(samplesPerPixel)
	r = math.Sqrt(scale * r)
	g = math.Sqrt(scale * g)
	b = math.Sqrt(scale * b)

	// Write the translated [0,255] value of each color component.
	return color.RGBA{
		uint8(256 * clamp(r, 0.0, 0.999)),
		uint8(256 * clamp(g, 0.0, 0.999)),
		uint8(256 * clamp(b, 0.0, 0.999)),
		255,
	}
}
