package jsontag

import (
	"encoding/json"
	"testing"
)

func BenchmarkJsonTagSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var w WithTagSmall
		json.Unmarshal(rawjsonSmall, &w)
	}
}

func BenchmarkJsonNoTagSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var w WithoutTagSmall
		json.Unmarshal(rawjsonSmall, &w)
	}
}

func BenchmarkJsonTagMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var w WithTagMedium
		json.Unmarshal(rawjsonMedium, &w)
	}
}

func BenchmarkJsonNoTagMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var w WithoutTagMedium
		json.Unmarshal(rawjsonMedium, &w)
	}
}

func BenchmarkJsonTagLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var w WithTagLarge
		json.Unmarshal(rawjsonLarge, &w)
	}
}

func BenchmarkJsonNoTagLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var w WithoutTagLarge
		json.Unmarshal(rawjsonLarge, &w)
	}
}
