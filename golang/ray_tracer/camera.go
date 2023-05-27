package ray_tracer

import "math"

type Camera struct {
	Width, Height int
	Fov float64
	Transform Matrix
	pixelSize float64
	worldHfWidth, worldHfHeight float64
}

func NewCamera(width, height int, fov float64) Camera {
	halfView := math.Tan(fov / 2)
	aspect := float64(width) / float64(height)
	var worldHfHeight, worldHfWidth float64
	if aspect > 1.0 {
		worldHfWidth = halfView
		worldHfHeight = halfView / float64(aspect)
	} else {
		worldHfHeight = halfView
		worldHfWidth = halfView * aspect
	}
	pixelSize := worldHfWidth * 2 / float64(width)
	return Camera{width, height, fov, NewIdentityMatrix(4), pixelSize, worldHfWidth, worldHfHeight}
}

func (c *Camera) SetTransform(transform Matrix) *Camera {
	c.Transform = transform
	return c
}

func (c Camera) PixelSize() float64 {
	return c.pixelSize
}

func (c Camera) PixelRay(px, py int) Ray {
	offsetX := (float64(px) + 0.5) * c.pixelSize
	offsetY := (float64(py) + 0.5) * c.pixelSize

	worldX := c.worldHfWidth - offsetX
	worldY := c.worldHfHeight - offsetY

	invTransform, _ := c.Transform.Inverse()
	pixel, _ := invTransform.Multiply4(NewPoint(worldX, worldY, -1).Tuple()).Point()
	origin, _ := invTransform.Multiply4(NewPoint(0, 0, 0).Tuple()).Point()
	direction := pixel.Subtract(origin).Normalize()
	return NewRay(origin, direction)
}

func (c Camera) Render(w World) Canvas {
	image := NewCanvas(c.Width, c.Height)

	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			ray := c.PixelRay(x, y)
			image[y][x] = w.ColorAt(ray)
		}
	}
	return image
}
