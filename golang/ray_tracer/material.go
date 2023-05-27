package ray_tracer

import (
	"math"
)

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() Material {
	return Material{NewColor(1, 1, 1), 0.1, 0.9, 0.9, 200.0}
}

func (m *Material) SetColor(c Color) *Material {
	m.Color = c
	return m
}

func (m *Material) SetAmbient(a float64) *Material {
	m.Ambient = a
	return m
}

func (m *Material) SetDiffuse(d float64) *Material {
	m.Diffuse = d
	return m
}

func (m *Material) SetSpecular(s float64) *Material {
	m.Specular = s
	return m
}

func (m *Material) SetShininess(s float64) *Material {
	m.Shininess = s
	return m
}

func Lighting(m Material, p Point, l PointLight, ev, nv Vector) Color {
	effC := m.Color.Multiply(l.Intensity)
	lv := l.Position.Subtract(p).Normalize()
	amb := effC.MultiplyScalar(m.Ambient)
	dif := NewColor(0, 0, 0)
	spc := NewColor(0, 0, 0)

	if ldn := DotProduct(lv, nv); ldn >= 0 {
		dif = effC.MultiplyScalar(m.Diffuse * ldn)
		rv := lv.Negate().Reflect(nv)
		if rde := DotProduct(ev, rv); rde > 0 {
			var f float64;
			if almostEqual(1.0, roundFloat(rde, 5)) {
				f = 1.0
			} else {
				f = roundFloat(math.Pow(rde, m.Shininess), 4)
			}
			spc = l.Intensity.MultiplyScalar(m.Specular * f)
		}
	}
	return amb.Add(dif).Add(spc)
}
