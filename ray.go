package main

import (
	"image/color"
	"math"
)

type Ray struct {
	origin    Vec3
	direction Vec3
}

func (r Ray) At(t float64) Vec3 {
	v := r.direction.ScalarMul(t)
	return r.origin.Add(v)
}

func (r Ray) Color() color.RGBA {
	t := r.hitSphere(Vec3{0, 0, -1}, 0.5)
	if t > 0.0 {
		N := r.At(t).Unit().Sub(Vec3{0, 0, -1})
		return Vec3{N.x + 1, N.y + 1, N.z + 1}.ScalarMul(0.5).RGBA()
	}
	unitDirection := r.direction.Unit()
	t = 0.5 * (unitDirection.y + 1.0)
	a := Vec3{1.0, 1.0, 1.0}.ScalarMul(1.0 - t)
	b := Vec3{0.5, 0.7, 1.0}.ScalarMul(t)
	return a.Add(b).RGBA()

}

func (r Ray) hitSphere(center Vec3, radius float64) float64 {
	oc := r.origin.Sub(center)
	a := r.direction.Dot(r.direction)
	b := 2.0 * oc.Dot(r.direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	}
	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}
