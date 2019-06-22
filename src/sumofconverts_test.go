package largeoutput

import (
	"fmt"
	"testing"
)

/*
Recursive is marginally better

go test -bench=SumOfConverts
goos: windows
goarch: amd64
BenchmarkSumOfConverts-4          500000              2539 ns/op
BenchmarkSumOfConvertsR-4         500000              2493 ns/op
PASS
*/
func ExampleSumOfConverts() {
	fmt.Println(sumOfConverts()) // fmt is required to check // Output:
	// Output: 14
}

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
