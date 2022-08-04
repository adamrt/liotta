package main

type Material interface {
	Scatter(ray Ray, record *HitRecord, attenuation *Vec3, scattered *Ray) bool
}

type Lambert struct {
	albedo Vec3
}

func (l Lambert) Scatter(ray Ray, record *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	var scatterDirection Vec3
	switch LambertDiffuseMethod {
	case 1:
		scatterDirection = record.normal.Add(Vec3RandomInUnitSphere())
	case 2:
		scatterDirection = record.normal.Add(Vec3RandUnit())
	case 3:
		scatterDirection = Vec3RandomInHemisphere(record.normal)
	}

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = record.normal
	}

	*scattered = Ray{record.point, scatterDirection}
	*attenuation = l.albedo
	return true
}

type Metal struct {
	albedo Vec3
}

func (m Metal) Scatter(ray Ray, record *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	reflected := ray.direction.Unit().Reflect(record.normal)
	*scattered = Ray{record.point, reflected}
	*attenuation = m.albedo
	return scattered.direction.Dot(record.normal) > 0
}

// ray := Ray{record.point, target.Sub(record.point)}
// return ray.Color(world, depth-1).ScalarMul(0.5)
