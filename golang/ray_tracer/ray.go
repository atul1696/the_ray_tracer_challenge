package ray_tracer

type Ray struct {
	Origin    Point
	Direction Vector
}

func NewRay(p Point, v Vector) Ray {
	return Ray{p, v}
}

func (r Ray) Equals(r1 Ray) bool {
	return r.Origin.Equals(r1.Origin) && r.Direction.Equals(r1.Direction)
}

func (r Ray) Position(t float64) Point {
	return r.Origin.AddVector(r.Direction.MultiplyScalar(t))
}

func (r Ray) Transform(t Matrix) (Ray, error) {
	o, err1 := r.Origin.Transform(t)
	d, err2 := r.Direction.Transform(t)
	if err1 != nil || err2 != nil {
		return Ray{}, err1
	}
	return NewRay(o, d), nil
}

// Assumes that the given matrix is 4x4
func (r Ray) Transform4(t Matrix) Ray {
	o, _ := r.Origin.Transform(t)
	d, _ := r.Direction.Transform(t)
	return NewRay(o, d)
}
