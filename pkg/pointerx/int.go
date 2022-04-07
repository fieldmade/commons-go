package pointerx

func Int(s int) *int {
	return &s
}

func IntR(s *int) int {
	if s == nil {
		return int(0)
	}
	return *s
}

func Int32(s int32) *int32 {
	return &s
}

func Int32R(s *int32) int32 {
	if s == nil {
		return int32(0)
	}
	return *s
}

func Int64(s int64) *int64 {
	return &s
}

func Int64R(s *int64) int64 {
	if s == nil {
		return int64(0)
	}
	return *s
}
