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

func Lighting(m Material, p Point, light PointLight, eyeVec, normalVec Vector, inShadow bool) Color {
	effectiveColor := m.Color.Multiply(light.Intensity)
	ambient := effectiveColor.MultiplyScalar(m.Ambient)

	if (inShadow) {
		return ambient
	}

	diffuse := NewColor(0, 0, 0)
	specular := NewColor(0, 0, 0)

	lightVec := light.Position.Subtract(p).Normalize()
	if lightDotNormal := DotProduct(lightVec, normalVec); lightDotNormal >= 0 {
		diffuse = effectiveColor.MultiplyScalar(m.Diffuse * lightDotNormal)
		reflectVec := lightVec.Negate().Reflect(normalVec)
		if reflectDotEye := DotProduct(eyeVec, reflectVec); reflectDotEye > 0 {
			var factor float64;
			if almostEqual(1.0, roundFloat(reflectDotEye, 5)) {
				factor = 1.0
			} else {
				factor = roundFloat(math.Pow(reflectDotEye, m.Shininess), 4)
			}
			specular = light.Intensity.MultiplyScalar(m.Specular * factor)
		}
	}
	return ambient.Add(diffuse).Add(specular)
}
