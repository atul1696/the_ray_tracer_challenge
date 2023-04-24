package main

import (
	"fmt"
	rt "ray_tracer"
)

func chapter5() {
	color := rt.NewColor(1, 0, 0)
	canvasSize := 100
	canvas := rt.NewCanvas(canvasSize, canvasSize)
	rayOrigin := rt.NewPoint(0, 0, -5)
	origin := rt.NewPoint(0, 0, 0)
	z := 10.0
	wallSize := 7.0
	translate := -wallSize / 2
	pixelSize := wallSize / float64(canvasSize)
	s := rt.NewSphere()

	// s.SetTransform(rt.Scale(1, 0.5, 1))
	// s.SetTransform(rt.Scale(0.5, 1, 1))
	// if t, err := rt.RotateZ(math.Pi / 4).Multiply(rt.Scale(0.5, 1, 1)); err == nil {
	// 	s.SetTransform(t)
	// }
	// s.SetTransform(rt.Shear(1, 0, 0, 0, 0, 0))

	for i := 0; i < canvasSize; i++ {
		x := float64(i)*pixelSize + translate
		for j := 0; j < canvasSize; j++ {
			y := float64(j)*pixelSize + translate
			point := rt.NewPoint(x, y, z)
			rayDirection := point.Subtract(origin).Normalize()
			ray := rt.NewRay(rayOrigin, rayDirection)
			xs := s.Intersect(ray)
			if _, ok := xs.Hit(); ok {
				canvas[i][j] = color
			}
		}
	}
	fmt.Print(canvas.CreatePpm())
}
