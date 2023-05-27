package ray_tracer

import "sort"

type Intersection struct {
	T     float64
	Shape Sphere
}

func NewIntersection(t float64, o Sphere) Intersection {
	return Intersection{t, o}
}

func (x Intersection) Equals(x1 Intersection) bool {
	return almostEqual(x.T, x1.T) && x.Shape.Equals(x1.Shape)
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
	xs.sort()
	i := sort.Search(len(xs), func(i int) bool { return xs[i].T >= 0 })
	if i >= len(xs) {
		return Intersection{}, false
	}
	return xs[i], true
}

func (xs *Intersections) sort() {
	sort.Slice(*xs, func(i, j int) bool {
		return (*xs)[i].T < (*xs)[j].T
	})
}
