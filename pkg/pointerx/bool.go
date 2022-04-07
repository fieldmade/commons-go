package pointerx

func Bool(s bool) *bool {
	return &s
}

func BoolR(s *bool) bool {
	if s == nil {
		return false
	}
	return *s
}
