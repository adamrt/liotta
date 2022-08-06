package main

import "math"

func NewSphere(center Vec3, radius float64, material Material) Sphere {
	return Sphere{center: center, radius: radius, material: material}
}

type Sphere struct {
	center   Vec3
	radius   float64
	material Material
}

func (s Sphere) Hit(ray *Ray, tMin, tMax float64) (bool, *HitRecord) {
	oc := ray.origin.Sub(s.center)
	a := ray.direction.Dot(ray.direction)
	halfB := oc.Dot(ray.direction)
	c := oc.Dot(oc) - s.radius*s.radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false, nil
	}
	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (-halfB - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtd) / a
		if root < tMin || tMax < root {
			return false, nil
		}
	}

	point := ray.At(root)
	record := HitRecord{
		t:        root,
		point:    point,
		material: s.material,
	}
	outwardNormal := point.Sub(s.center).ScalarDiv(s.radius)
	record.SetFaceNormal(ray, outwardNormal)

	return true, &record
}
