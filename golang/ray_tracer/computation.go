package ray_tracer

type Computation struct {
	T float64
	Object Sphere
	Point Point
	EyeV Vector
	NormalV Vector
	Inside bool
}

func PrepareComputation(x Intersection, r Ray) Computation {
	p := r.Position(x.T)
	eyeV := r.Direction.Negate()
	normalV := x.Shape.NormatAt(p)
	inside := false
	if DotProduct(eyeV, normalV) < 0 {
		inside = true
		normalV = normalV.Negate()
	}
	return Computation{
		x.T,
		x.Shape,
		p,
		eyeV,
		normalV,
		inside,
	}
}
