package main

import (
	"fmt"
	rt "ray_tracer"
)

func chapter6() {
	canvasSize := 100
	canvas := rt.NewCanvas(canvasSize, canvasSize)
	rayOrigin := rt.NewPoint(0, 0, -5)
	origin := rt.NewPoint(0, 0, 0)
	z := 10.0
	wallSize := 7.0
	translate := wallSize / 2
	pixelSize := wallSize / float64(canvasSize)
	mat := rt.NewMaterial()
	mat.SetColor(rt.NewColor(1, 0.2, 1))
	s := rt.NewSphere()
	s.SetMaterial(mat)

	s.SetTransform(rt.Scale(1.5, 1.5, 1.5))
	// s.SetTransform(rt.Scale(0.5, 1, 1))
	// if t, err := rt.RotateZ(math.Pi / 4).Multiply(rt.Scale(0.5, 1, 1)); err == nil {
	// 	s.SetTransform(t)
	// }
	// s.SetTransform(rt.Shear(1, 0, 0, 0, 0, 0))

	light := rt.NewPointLight(rt.NewPoint(-10, 10, -10), rt.NewColor(1, 1, 1))

	for i := 0; i < canvasSize; i++ {
		y := translate - float64(i)*pixelSize
		for j := 0; j < canvasSize; j++ {
			x := -translate + float64(j)*pixelSize
			point := rt.NewPoint(x, y, z)
			rayDirection := point.Subtract(origin).Normalize()
			ray := rt.NewRay(rayOrigin, rayDirection)
			xs := s.Intersect(ray)
			if x, ok := xs.Hit(); ok {
				p := ray.Position(x.T)
				ev := rayDirection.Negate()
				nv := x.Shape.NormatAt(p)
				c := rt.Lighting(x.Shape.Material, p, light, ev, nv)
				canvas[i][j] = c
			}
		}
	}
	fmt.Print(canvas.CreatePpm())
}
