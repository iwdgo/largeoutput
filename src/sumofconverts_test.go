package largeoutput

import (
	"testing"
)

/*
No significant difference between recursive and non-recursive.

go version go1.13.1 windows/amd64
>go test -bench=SumOfConverts
goos: windows
goarch: amd64
BenchmarkSumOfConverts-4          333518              3419 ns/op
BenchmarkSumOfConvertsR-4         343140              3489 ns/op
PASS
*/

func TestSumOfConverts(t *testing.T) {
	got := sumOfConverts()
	want := int64(14)
	if got != want {
		t.Errorf("sum of converts : got %d, want %d", got, want)
	}
}

/* */
func BenchmarkSumOfConverts(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		sumOfConverts()
	}
}

// Recursive
func TestSumOfConvertsR(t *testing.T) {
	got := sumOfConvertsR()
	want := int64(14)
	if got != want {
		t.Errorf("sum of converts recursive : got %d, want %d", got, want)
	}
}

func BenchmarkSumOfConvertsR(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		sumOfConvertsR()
	}
}
