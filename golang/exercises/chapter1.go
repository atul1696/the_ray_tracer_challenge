package main

import (
	"fmt"
	rt "ray_tracer"
)

func chapter1() {
	proj := Projectile{rt.NewPoint(0, 1, 0), rt.NewVector(1, 1, 0).Normalize()}
	env := Environment{rt.NewVector(0, -0.1, 0), rt.NewVector(-0.01, 0, 0)}

	var tickCount int
	for ; proj.position.Y > 0; tickCount++ {
		fmt.Println(proj)
		proj = tick(env, proj)
	}
	fmt.Printf("Tick count = %d\n", tickCount)
}
