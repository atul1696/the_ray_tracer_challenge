package ray_tracer

import "math"

const float64EqualityThreshold = 1e-5

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func limit(v float64, min, max int) int {
	val := int(v)
	if val < min {
		val = min
	} else if val > max {
		val = max
	}
	return val
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
