package main

import "math"

func NewCamera(
	position Vec3,
	lookAt Vec3,
	up Vec3,
	vfov,
	aspectRatio float64,
	aperture float64,
	focusDist float64,
) Camera {

	theta := degreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := position.Sub(lookAt).Unit()
	u := up.Cross(w).Unit()
	v := w.Cross(u)

	origin := position
	horizontal := u.ScalarMul(focusDist * viewportWidth)
	vertical := v.ScalarMul(focusDist * viewportHeight)
	lowerLeftCorner := origin.Sub(horizontal.ScalarDiv(2)).Sub(vertical.ScalarDiv(2)).Sub(w.ScalarMul(focusDist))
	lensRadius := aperture / 2

	return Camera{
		origin:          origin,
		horizontal:      horizontal,
		vertical:        vertical,
		lensRadius:      lensRadius,
		lowerLeftCorner: lowerLeftCorner,
		u:               u,
		v:               v,
		w:               w,
	}
}

type Camera struct {
	origin          Vec3
	horizontal      Vec3
	vertical        Vec3
	lowerLeftCorner Vec3
	u, v, w         Vec3
	lensRadius      float64
}

func (c Camera) GetRay(s, t float64) Ray {
	rd := Vec3RandomInUnitDisk().ScalarMul(c.lensRadius)
	offset := c.u.ScalarMul(rd.x).Add(c.v.ScalarMul(rd.y))

	direction := c.lowerLeftCorner.Add(c.horizontal.ScalarMul(s)).Add(c.vertical.ScalarMul(t)).Sub(c.origin).Sub(offset)

	return Ray{
		origin:    c.origin.Add(offset),
		direction: direction,
	}
}
