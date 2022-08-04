package main

import "math"

type Material interface {
	Scatter(ray Ray, record *HitRecord, attenuation *Vec3, scattered *Ray) bool
}

func NewLambert(albedo Vec3) Lambert {
	return Lambert{albedo: albedo}
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

func NewMetal(albedo Vec3, fuzz float64) Metal {
	return Metal{albedo: albedo, fuzz: fuzz}
}

type Metal struct {
	albedo Vec3
	fuzz   float64
}

func (m Metal) Scatter(ray Ray, record *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	reflected := ray.direction.Unit().Reflect(record.normal)
	*scattered = Ray{record.point, Vec3RandomInUnitSphere().ScalarMul(m.fuzz).Add(reflected)}
	*attenuation = m.albedo
	return scattered.direction.Dot(record.normal) > 0
}

func NewDielectric(index_of_refraction float64) Dielectric {
	return Dielectric{ir: index_of_refraction}
}

type Dielectric struct {
	// Index of Refraction
	ir float64
}

func (d Dielectric) Scatter(ray Ray, record *HitRecord, attenuation *Vec3, scattered *Ray) bool {

	refractionRatio := d.ir
	if record.frontFace {
		refractionRatio = 1.0 / d.ir
	}

	unitDirection := ray.direction.Unit()

	cosTheta := math.Min(unitDirection.Neg().Dot(record.normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1.0
	direction := Vec3{}

	if cannotRefract || d.reflectance(cosTheta, refractionRatio) > random() {
		direction = unitDirection.Reflect(record.normal)
	} else {
		direction = unitDirection.Refract(record.normal, refractionRatio)
	}

	*scattered = Ray{record.point, direction}

	*attenuation = Vec3{1.0, 1.0, 1.0}
	return true

}

// Use Schlick's approximation for reflectance.
func (d Dielectric) reflectance(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
