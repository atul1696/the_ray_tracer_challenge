package ray_tracer

import "math"

type Sphere struct {
	TransformMat Matrix
}

func (s Sphere) Equals(s1 Sphere) bool {
	if s.TransformMat.Rows() != s1.TransformMat.Rows() || s.TransformMat.Columns() != s1.TransformMat.Columns() {
		return false
	}
	for i := 0; i < s.TransformMat.Rows(); i++ {
		for j := 0; j < s.TransformMat.Columns(); j++ {
			if !almostEqual(s.TransformMat[i][j], s1.TransformMat[i][j]) {
				return false
			}
		}
	}
	return true
}

func NewSphere() Sphere {
	return Sphere{NewIdentityMatrix(4)}
}

func (s *Sphere) SetTransform(m Matrix) {
	(*s).TransformMat = m
}

func (s Sphere) Intersect(r Ray) Intersections {
	if t, err := s.TransformMat.Inverse(); err == nil {
		r = r.Transform4(t)
	}
	xs := make([]Intersection, 0, 2)
	sRay := r.Origin.Subtract(NewPoint(0, 0, 0))
	a := DotProduct(r.Direction, r.Direction)
	b := 2 * DotProduct(r.Direction, sRay)
	c := DotProduct(sRay, sRay) - 1.0

	discriminant := b*b - 4*a*c
	if discriminant >= 0 {
		var t float64
		t = (-b - math.Sqrt(discriminant)) / (2 * a)
		xs = append(xs, NewIntersection(t, s))
		t = (-b + math.Sqrt(discriminant)) / (2 * a)
		xs = append(xs, NewIntersection(t, s))
	}
	return xs
}
