package largeoutput

import (
	"testing"
)

/*
File operation is the most consuming. One file less means half the time.

goos: windows
goarch: amd64
BenchmarkGetHTMLPageString-4                           1        1145923100 ns/op
BenchmarkGetHTMLPageBuffer-4                           1        1030399200 ns/op
BenchmarkGetHTMLPageBufferNoGotFile-4                  2         582395500 ns/op
*/

/* go test -run=TestGetHTMLPageString */
func TestGetHTMLPageString(t *testing.T) {
	GetHTMLPageString()
}

/* go test -bench=GetHTMLPageString */
func BenchmarkGetHTMLPageString(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHTMLPageString()
	}
}

/* go test -run=TestGetHTMLPageBuffer */
func TestGetHTMLPageBuffer(t *testing.T) {
	GetHTMLPageBuffer()
}

/* go test -bench=GetHTMLPageBuffer */
func BenchmarkGetHTMLPageBuffer(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHTMLPageBuffer()
	}
}

func TestGetHTMLPageBufferNoGotFile(t *testing.T) {
	GetHTMLPageBufferNoGotFile()
}

/* go test -bench=GetHTMLPageBuffer */
func BenchmarkGetHTMLPageBufferNoGotFile(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHTMLPageBufferNoGotFile()
	}
}
