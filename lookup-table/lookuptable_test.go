package lookuptable

import (
	"fmt"
	"io"
	"testing"
)

func BenchmarkValidProfilingType(b *testing.B) {
	blackhole := false
	for i := 0; i < b.N; i++ {
		blackhole = ValidProfilingType("go", "cpu")
	}
	fmt.Fprint(io.Discard, blackhole)
}

func BenchmarkValidProfilingTypeLookup(b *testing.B) {
	blackhole := false
	for i := 0; i < b.N; i++ {
		blackhole = ValidProfilingTypeLookup("go", "cpu")
	}
	fmt.Fprint(io.Discard, blackhole)
}

func BenchmarkValidProfilingTypeNestedLookup(b *testing.B) {
	blackhole := false
	for i := 0; i < b.N; i++ {
		blackhole = ValidProfilingTypeNestedLookup("go", "cpu")
	}
	fmt.Fprint(io.Discard, blackhole)
}

func BenchmarkValidProfilingTypeArrays(b *testing.B) {
	blackhole := false
	for i := 0; i < b.N; i++ {
		blackhole = ValidProfilingTypeArrays("go", "cpu")
	}
	fmt.Fprint(io.Discard, blackhole)
}
