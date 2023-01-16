package sumofconverts

import (
	"testing"
)

var benchstrings = []string{"a", "b", "c", "4", "e", "f", "g", "8", "i", "j", "k"}

func BenchmarkSumOfConverts(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumOfConverts(benchstrings)
	}
}

func BenchmarkSumOfConvertsR(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumOfConvertsR(benchstrings)
	}
}

func BenchmarkSumOfConvertsRPtr(b *testing.B) {
	bcopy := benchstrings
	for n := 0; n < b.N; n++ {
		SumOfConvertsRPtr(&bcopy)
	}
}
