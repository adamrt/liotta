package main

func NewCamera() Camera {
	viewportHeight := 2.0
	viewportWidth := AspectRatio * viewportHeight
	focalLength := 1.0

	origin := Vec3{0, 0, 0}
	horizontal := Vec3{viewportWidth, 0, 0}
	vertical := Vec3{0, viewportHeight, 0}
	lowerLeftCorner := origin.Sub(horizontal.ScalarDiv(2)).Sub(vertical.ScalarDiv(2)).Sub(Vec3{0, 0, focalLength})
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
