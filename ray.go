package main

import "image/color"

type Ray struct {
	origin    Vec3
	direction Vec3
}

func (r Ray) At(t float64) Vec3 {
	v := r.direction.ScalarMul(t)
	return r.origin.Add(v)
}

func (r Ray) Color() color.RGBA {
	unitDirection := r.direction.Unit()
	t := 0.5 * (unitDirection.y + 1.0)
	a := Vec3{1.0, 1.0, 1.0}.ScalarMul(1.0 - t)
	b := Vec3{0.5, 0.7, 1.0}.ScalarMul(t)
	return a.Add(b).RGBA()
}
