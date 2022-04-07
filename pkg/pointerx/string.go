package pointerx

func String(s string) *string {
	return &s
}

func StringR(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
