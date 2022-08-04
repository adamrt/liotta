package main

import "math"

type Sphere struct {
	center Vec3
	radius float64
}

func (s Sphere) Hit(ray Ray, tMin, tMax float64) (bool, HitRecord) {
	oc := ray.origin.Sub(s.center)
	a := ray.direction.LengthSquared()
	halfB := oc.Dot(ray.direction)
	c := oc.LengthSquared() - s.radius*s.radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false, HitRecord{}
	}
	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (-halfB - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtd) / a
		if root < tMin || tMax < root {
			return false, HitRecord{}
		}
	}

	point := ray.At(root)
	record := HitRecord{
		t:     root,
		point: point,
	}
	outwardNormal := point.Sub(s.center).ScalarDiv(s.radius)
	record.SetFaceNormal(ray, outwardNormal)

	return true, record
}
