package main

import rt "ray_tracer"

type Projectile struct {
	position rt.Point
	velocity rt.Vector
}

type Environment struct {
	gravity rt.Vector
	wind    rt.Vector
}

func tick(env Environment, proj Projectile) Projectile {
	pos := proj.position.AddVector(proj.velocity)
	vel := proj.velocity.Add(env.gravity).Add(env.wind)
	return Projectile{pos, vel}
}

func limit(v float64, min, max int) int {
	val := int(v)
	if val < min {
		val = min
	} else if val > max {
		val = max
	}
	return val
}
