package main

// Hittable is an interface representing something that can be hit by rays.  It
// can be a single item like a Sphere or a group of objects like the World
// (array of Spheres).
type Hittable interface {
	Hit(ray Ray, min, max float64, record *HitRecord) bool
}

// Hit Record stores information about a specific intersect of a ray hitting a
// Hittable object.
type HitRecord struct {
	point     Vec3
	normal    Vec3
	t         float64
	frontFace bool
}

func (r *HitRecord) SetFaceNormal(ray Ray, outwardNormal Vec3) {
	r.frontFace = ray.direction.Dot(outwardNormal) < 0
	if r.frontFace {
		r.normal = outwardNormal
	} else {
		r.normal = outwardNormal.Neg()
	}
}

// World is an array of hittable objects (Spheres, etc) that implement the
// Hittable interface. It handles the z-index to know what to draw in front.

type World []Hittable

func (w World) Hit(ray Ray, tMin, tMax float64, record *HitRecord) bool {
	tempRecord := HitRecord{}
	hitAnything := false
	closestSoFar := tMax

	for _, object := range w {
		if object.Hit(ray, tMin, closestSoFar, &tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.t
			*record = tempRecord
		}
	}
	return hitAnything
}
