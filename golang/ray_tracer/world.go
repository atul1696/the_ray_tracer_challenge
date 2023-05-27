package ray_tracer

type World struct {
	Shapes []Sphere
	Light  PointLight
}

func NewWorld() World {
	return World{}
}

func NewDefaultWorld() World {
	s1 := NewSphere()
	m := NewMaterial()
	m.SetColor(NewColor(0.8, 1.0, 0.6)).SetDiffuse(0.7).SetSpecular(0.2)
	s1.SetMaterial(m)
	s2 := NewSphere()
	s2.SetTransform(Scale(0.5, 0.5, 0.5))
	return World{[]Sphere{s1, s2}, NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1))}
}

func (w *World) SetLight(l PointLight) *World {
	w.Light = l
	return w
}

func (w *World) AddShape(s Sphere) *World {
	w.Shapes = append(w.Shapes, s)
	return w
}

func (w *World) AddShapes(shapes... Sphere) *World {
	for _, shape := range shapes {
		w.Shapes = append(w.Shapes, shape)
	}
	return w
}

func (w *World) RemoveShapes() *World {
	w.Shapes = make([]Sphere, 0)
	return w
}

func (w World) Intersect(r Ray) Intersections {
	var xs Intersections = make([]Intersection, 0)
	for _, shape := range w.Shapes {
		for _, x := range shape.Intersect(r) {
			xs.append(x)
		}
	}
	xs.sort()
	return xs
}

func (w World) ShadeHit(c Computation) Color {
	return Lighting(c.Object.Material, c.Point, w.Light, c.EyeV, c.NormalV, w.isInShadow(c.overPoint))
}

func (w World) ColorAt(r Ray) Color {
	xs := w.Intersect(r)
	x, ok := xs.Hit()
	if !ok {
		return NewColor(0, 0, 0)
	}
	return w.ShadeHit(PrepareComputation(x, r))
}

func (w World) isInShadow(p Point) bool {
	dir := w.Light.Position.Subtract(p)
	dist := dir.Magnitude()
	r := NewRay(p, dir.Normalize())
	xs := w.Intersect(r)
	x, has := xs.Hit()
	return has && x.T < dist
}
