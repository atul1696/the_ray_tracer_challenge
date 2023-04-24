package ray_tracer

import "sort"

type Intersection struct {
	t float64
	s Sphere
}

func NewIntersection(t float64, o Sphere) Intersection {
	return Intersection{t, o}
}

func (x Intersection) Equals(x1 Intersection) bool {
	return almostEqual(x.t, x1.t) && x.s.Equals(x1.s)
}

type Intersections []Intersection

func NewIntersections(is ...Intersection) Intersections {
	xs := make([]Intersection, 0)
	xs = append(xs, is...)
	return xs
}

func (xs *Intersections) append(x Intersection) {
	*xs = append(*xs, x)
}

func (xs Intersections) Hit() (Intersection, bool) {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].t < xs[j].t
	})
	i := sort.Search(len(xs), func(i int) bool { return xs[i].t >= 0 })
	if i >= len(xs) {
		return Intersection{}, false
	}
	return xs[i], true
}
