package asymptoticcomplexity

// CreateMap creates a map with O(1) complexity
func CreateMap(size int) map[int]bool {
	m := make(map[int]bool, size)
	for i := 0; i < size; i++ {
		m[i] = true
	}
	return m
}

// CreateSlice creates a slice with O(1) complexity
func CreateSlice(size int) []bool {
	s := make([]bool, size)
	for i := 0; i < size; i++ {
		s[i] = true
	}
	return s
}
