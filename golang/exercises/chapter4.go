package main

import (
	"math"
	"fmt"
	rt "ray_tracer"
)

func chapter4() {
	point := rt.NewPoint(120, 0, 0)
	canvas := rt.NewCanvas(300, 300)
	color := rt.NewColor(0, 1, 0)

	rot := rt.RotateY(math.Pi / 6)
	// tls := rt.Translate(150, 0, 150)

	for t := 0; t < 12; t++ {
		// p, _ := point.Transform(tls)
		i, j := int(point.X) + 150, int(point.Z) + 150
		canvas[i][j] = color
		point, _ = point.Transform(rot)
	}

	fmt.Print(canvas.CreatePpm())
}
