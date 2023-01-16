package sumofconverts

import (
	"strconv"
	"testing"
)

func TestEndWithNumber(t *testing.T) {
	var want int64 = 2
	a := []string{strconv.Itoa(int(want))}
	if got := SumOfConverts(a); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got := SumOfConvertsR(a); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got := SumOfConvertsRPtr(&a); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestEndWithString(t *testing.T) {
	var want int64 = 0
	a := []string{"a"}
	if got := SumOfConverts(a); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got := SumOfConvertsR(a); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got := SumOfConvertsRPtr(&a); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

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
