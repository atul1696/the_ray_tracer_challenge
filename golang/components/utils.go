package components

import "math"

const float64EqualityThreshold = 1e-5

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
