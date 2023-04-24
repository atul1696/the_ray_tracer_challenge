package ray_tracer

func sliceEquals(s1, s2 []float64) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func floatSlice(xs Intersections) []float64 {
	fs := make([]float64, 0)
	for _, x := range xs {
		fs = append(fs, x.t)
	}
	return fs
}
