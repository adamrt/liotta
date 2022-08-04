package main

import (
	"math"
)

var Infinity = math.Inf(1)

type Ray struct {
	origin    Vec3
	direction Vec3
}

func (r Ray) At(t float64) Vec3 {
	v := r.direction.ScalarMul(t)
	return r.origin.Add(v)
}

func (r Ray) Color(world Hittable, depth int) Vec3 {
	// If we've exceeded the ray bounce limit, no more light is gathered.
	if depth <= 0 {
		return Vec3{0, 0, 0}
	}

	hit, record := world.Hit(r, 0.001, Infinity)
	if hit {
		target := record.point.Add(record.normal).Add(Vec3RandomInUnitSphere())
		ray := Ray{record.point, target.Sub(record.point)}
		return ray.Color(world, depth-1).ScalarMul(0.5)
	}

	unitDirection := r.direction.Unit()
	t := 0.5 * (unitDirection.y + 1.0)
	a := Vec3{1.0, 1.0, 1.0}.ScalarMul(1.0 - t)
	b := Vec3{0.5, 0.7, 1.0}.ScalarMul(t)
	return a.Add(b)

}

func (r Ray) hitSphere(center Vec3, radius float64) float64 {
	oc := r.origin.Sub(center)
	a := r.direction.LengthSquared()
	halfB := oc.Dot(r.direction)
	c := oc.LengthSquared() - radius*radius
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return -1.0
	}
	return (-halfB - math.Sqrt(discriminant)) / a
}
