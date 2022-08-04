package main

import "math"

func NewCamera(position Vec3, lookAt Vec3, up Vec3, vfov, aspectRatio float64) Camera {

	theta := degreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := position.Sub(lookAt).Unit()
	u := up.Cross(w).Unit()
	v := w.Cross(u)

	origin := position
	horizontal := u.ScalarMul(viewportWidth)
	vertical := v.ScalarMul(viewportHeight)
	lowerLeftCorner := origin.Sub(horizontal.ScalarDiv(2)).Sub(vertical.ScalarDiv(2)).Sub(w)

	return Camera{
		origin:          origin,
		horizontal:      horizontal,
		vertical:        vertical,
		lowerLeftCorner: lowerLeftCorner,
	}
}

type Camera struct {
	origin          Vec3
	horizontal      Vec3
	vertical        Vec3
	lowerLeftCorner Vec3
}

func (c Camera) GetRay(u, v float64) Ray {
	direction := c.lowerLeftCorner.Add(c.horizontal.ScalarMul(u)).Add(c.vertical.ScalarMul(v)).Sub(c.origin)
	return Ray{c.origin, direction}
}
