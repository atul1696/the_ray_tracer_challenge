package ray_tracer

import "math"

func Translate(x, y, z float64) Matrix {
	m := NewIdentityMatrix(4)
	m[0][3] = x
	m[1][3] = y
	m[2][3] = z
	return m
}

func Scale(x, y, z float64) Matrix {
	m := NewIdentityMatrix(4)
	m[0][0] = x
	m[1][1] = y
	m[2][2] = z
	return m
}

func RotateX(r float64) Matrix {
	m := NewMatrix(4, 4)
	m[0][0] = 1.0
	m[1][1] = math.Cos(r)
	m[1][2] = -math.Sin(r)
	m[2][1] = math.Sin(r)
	m[2][2] = math.Cos(r)
	m[3][3] = 1.0
	return m
}

func RotateY(r float64) Matrix {
	m := NewMatrix(4, 4)
	m[0][0] = math.Cos(r)
	m[0][2] = math.Sin(r)
	m[1][1] = 1.0
	m[2][0] = -math.Sin(r)
	m[2][2] = math.Cos(r)
	m[3][3] = 1.0
	return m
}

func RotateZ(r float64) Matrix {
	m := NewMatrix(4, 4)
	m[0][0] = math.Cos(r)
	m[0][1] = -math.Sin(r)
	m[1][0] = math.Sin(r)
	m[1][1] = math.Cos(r)
	m[2][2] = 1.0
	m[3][3] = 1.0
	return m
}

func Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	m := NewIdentityMatrix(4)
	m[0][1] = xy
	m[0][2] = xz
	m[1][0] = yx
	m[1][2] = yz
	m[2][0] = zx
	m[2][1] = zy
	return m
}

func ViewTransform(from, to Point, up Vector) Matrix {
	forward := to.Subtract(from).Normalize()
	upNormalized := up.Normalize()
	left := CrossProduct(forward, upNormalized)
	trueUp := CrossProduct(left, forward)
	transform := NewMatrix(4, 4)
	transform[0][0] = left.X
	transform[0][1] = left.Y
	transform[0][2] = left.Z
	transform[1][0] = trueUp.X
	transform[1][1] = trueUp.Y
	transform[1][2] = trueUp.Z
	transform[2][0] = -forward.X
	transform[2][1] = -forward.Y
	transform[2][2] = -forward.Z
	transform[3][3] = 1.0

	transform, _ = transform.Multiply(Translate(-from.X, -from.Y, -from.Z))
	return transform
}
