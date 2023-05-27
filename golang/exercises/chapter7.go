package main

import (
	"fmt"
	"math"
	rt "ray_tracer"
)

func chapter7() {
	floor := rt.NewSphere()
	floor.SetTransform(rt.Scale(10, 0.01, 10))
	mt := rt.NewMaterial()
	mt.SetSpecular(0).SetColor(rt.NewColor(1, 0.9, 0.9))
	floor.SetMaterial(mt)

	leftWall := rt.NewSphere()
	leftWall.SetMaterial(mt)
	leftWall.SetTransform(rt.Translate(0, 0, 5).Multiply4(rt.RotateY(-math.Pi/4)).Multiply4(rt.RotateX(math.Pi/2)).Multiply4(rt.Scale(10, 0.01, 10)))

	rightWall := rt.NewSphere()
	rightWall.SetMaterial(mt)
	rightWall.SetTransform(rt.Translate(0, 0, 5).Multiply4(rt.RotateY(math.Pi/4)).Multiply4(rt.RotateX(math.Pi/2)).Multiply4(rt.Scale(10, 0.01, 10)))

	middleSphere := rt.NewSphere()
	middleSphere.SetTransform(rt.Translate(-0.5, 1, 0.5))
	mt = rt.NewMaterial()
	mt.SetColor(rt.NewColor(0.1, 1, 0.5)).SetDiffuse(0.7).SetSpecular(0.3)
	middleSphere.SetMaterial(mt)

	rightSphere := rt.NewSphere()
	rightSphere.SetTransform(rt.Translate(1.5, 0.5, -0.5).Multiply4(rt.Scale(0.5, 0.5, 0.5)))
	mt = rt.NewMaterial()
	mt.SetColor(rt.NewColor(0.5, 1, 0.1)).SetDiffuse(0.7).SetSpecular(0.3)
	rightSphere.SetMaterial(mt)

	leftSphere := rt.NewSphere()
	leftSphere.SetTransform(rt.Translate(-1.5, 0.33, -0.75).Multiply4(rt.Scale(0.33, 0.33, 0.33)))
	mt = rt.NewMaterial()
	mt.SetColor(rt.NewColor(1, 0.8, 0.1)).SetDiffuse(0.7).SetSpecular(0.3)
	leftSphere.SetMaterial(mt)

	world := rt.NewWorld()
	world.SetLight(rt.NewPointLight(rt.NewPoint(-10, 10, -10), rt.NewColor(1, 1, 1)))
	world.AddShapes(floor, leftWall, rightWall, middleSphere, rightSphere, leftSphere)

	camera := rt.NewCamera(100, 50, math.Pi/3)
	camera.SetTransform(rt.ViewTransform(rt.NewPoint(0, 1.5, -5), rt.NewPoint(0, 1, 0), rt.NewVector(0, 1, 0)))
	
	image := camera.Render(world)
	fmt.Println(image.CreatePpm())
}
