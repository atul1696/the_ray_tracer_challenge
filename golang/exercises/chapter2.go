package main

import (
	"fmt"
	rt "ray_tracer"
)

func chapter2() {
	color := rt.NewColor(0, 1, 0)
	width, height := 900, 550
	canvas := rt.NewCanvas(width, height)

	gravity := rt.NewVector(0, -0.1, 0)
	wind := rt.NewVector(-0.01, 0, 0)
	env := Environment{gravity, wind}

	pos := rt.NewPoint(0, 1, 0)
	vel := rt.NewVector(1, 1.8, 0).Normalize().MultiplyScalar(11.25)
	proj := Projectile{pos, vel}

	for proj.position.Y > 0 {
		x := limit(proj.position.X, 0, width-1)
		y := limit(proj.position.Y, 0, height-1)
		canvas[height-1-y][x] = color
		proj = tick(env, proj)
	}

	fmt.Print(canvas.CreatePpm())
}
