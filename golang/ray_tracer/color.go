package ray_tracer

type Color struct {
	R, G, B float64
}

func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

func (c Color) Add(c1 Color) Color {
	return NewColor(c.R+c1.R, c.G+c1.G, c.B+c1.B)
}

func (c Color) Subtract(c1 Color) Color {
	return NewColor(c.R-c1.R, c.G-c1.G, c.B-c1.B)
}

func (c Color) MultiplyScalar(factor float64) Color {
	return NewColor(c.R*factor, c.G*factor, c.B*factor)
}

func (c Color) Multiply(c1 Color) Color {
	return NewColor(c.R*c1.R, c.G*c1.G, c.B*c1.B)
}

func (c Color) Equals(c1 Color) bool {
	return almostEqual(c.R, c1.R) && almostEqual(c.G, c1.G) && almostEqual(c.B, c1.B)
}
