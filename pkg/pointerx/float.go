package pointerx

func Float32(s float32) *float32 {
	return &s
}

func Float32R(s *float32) float32 {
	if s == nil {
		return float32(0)
	}
	return *s
}

func Float64(s float64) *float64 {
	return &s
}

func Float64R(s *float64) float64 {
	if s == nil {
		return float64(0)
	}
	return *s
}
