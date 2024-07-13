package interfacestruct

import "testing"

func BenchmarkStructs(b *testing.B) {
	var stype []SomeInterface
	var sinter []SomeInterface
	var ptype []SomeType

	b.Run("Append to interface slice as type", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			stype = append(stype, SomeType{name: "test"})
		}
	})

	b.Run("Append to interface slice as interface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sinter = append(stype, SomeInterface(SomeType{name: "test"}))
		}
	})

	b.Run("Append to type slice as type", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ptype = append(ptype, SomeType{name: "test"})
		}
	})

	b.Run("Call method on interface (uncasted)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			stype[0].Some()
		}
	})

	b.Run("Call method on interface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sinter[0].Some()
		}
	})

	b.Run("Call method on type", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ptype[0].Some()
		}
	})

}
