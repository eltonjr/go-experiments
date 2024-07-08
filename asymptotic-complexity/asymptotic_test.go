package asymptoticcomplexity

import "testing"

func BenchmarkCreate(b *testing.B) {
	b.Run("Create map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = CreateMap(1000)
		}
	})

	b.Run("Create slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = CreateSlice(1000)
		}
	})
}

func BenchmarkAccess(b *testing.B) {
	s := CreateSlice(1000)
	m := CreateMap(1000)

	b.Run("Access map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = m[200]
		}
	})

	b.Run("Access slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = s[200]
		}
	})
}
