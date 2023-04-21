package ray_tracer

import (
	"fmt"
	"strings"
)

type Canvas [][]Color

func NewCanvas(width, height int) Canvas {
	c := make([][]Color, height)
	for i := range c {
		c[i] = make([]Color, width)
	}
	return c
}

func (c Canvas) Height() int {
	return len(c)
}

func (c Canvas) Width() int {
	return len(c[0])
}

func (c Canvas) CreatePpm() string {
	minColorValue, maxColorValue := 0, 255

	var sb strings.Builder
	sb.WriteString("P3\n")
	sb.WriteString(fmt.Sprintf("%d %d\n", c.Width(), c.Height()))
	sb.WriteString(fmt.Sprintln(maxColorValue))

	maxLineLength := 70
	for i := 0; i < c.Height(); i++ {
		lineLength := 0
		for j := 0; j < c.Width(); j++ {
			colorStr := ppmColor(c[i][j], minColorValue, maxColorValue)
			if lineLength+len(colorStr) >= maxLineLength {
				sb.WriteString("\n")
				lineLength = 0
			}
			sb.WriteString(colorStr)
			sb.WriteString(" ")
			lineLength += len(colorStr) + 1
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func ppmColor(c Color, min, max int) string {
	c = c.MultiplyScalar(float64(max))
	return fmt.Sprintf("%d %d %d", limit(c.R, min, max), limit(c.G, min, max), limit(c.B, min, max))
}
