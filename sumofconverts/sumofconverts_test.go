package sumofconverts

import (
	"testing"
)

/*
No significant difference between recursive and non-recursive.

goos: windows
goarch: amd64
pkg: github.com/iwdgo/largeoutput/sumofconverts
cpu: Intel(R) Core(TM) i5-5200U CPU @ 2.20GHz
BenchmarkSumOfConverts-4        60551074                94.29 ns/op
BenchmarkSumOfConvertsR-4       67506522                90.81 ns/op
BenchmarkSumOfConvertsRPtr-4    65597427                89.98 ns/op
PASS
ok      github.com/iwdgo/largeoutput/sumofconverts      18.090s
*/

func TestSumOfConverts(t *testing.T) {
	got := SumOfConverts(listofstrings)
	want := int64(total)
	if got != want {
		t.Errorf("sum of converts : got %d, want %d", got, want)
	}
}

func BenchmarkSumOfConverts(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		SumOfConverts(listofstrings)
	}
}

func TestSumOfConvertsR(t *testing.T) {
	got := SumOfConvertsR(listofstrings)
	want := int64(total)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func BenchmarkSumOfConvertsR(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumOfConvertsR(listofstrings)
	}
}

func TestSumOfConvertsRPtr(t *testing.T) {
	got := SumOfConvertsRPtr(&listofstrings)
	want := int64(total)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func BenchmarkSumOfConvertsRPtr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumOfConvertsRPtr(&listofstrings)
	}
}
