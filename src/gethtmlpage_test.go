// The first run of GetHTMLPage will fail and you can rename pagegot.html into pagewant.html
package largeoutput

import (
	"testing"
)

/*
File operation is the most consuming. One file less means half the time.
Buffer has a minor advantage over string.

goos: windows
goarch: amd64
BenchmarkGetHTMLPageString-4                           1        1059377800 ns/op
BenchmarkGetHTMLPageBuffer-4                           2         920314550 ns/op
BenchmarkGetHTMLPageBufferNoGotFile-4                  2         513047750 ns/op
*/

/* go test -run=TestGetHTMLPageString */
func TestGetHTMLPageString(t *testing.T) {
	if err := GetHTMLPageString(); err != nil {
		t.Error(err)
	}
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
	if err := GetHTMLPageBuffer(); err != nil {
		t.Error(err)
	}
}

/* go test -bench=GetHTMLPageBuffer */
func BenchmarkGetHTMLPageBuffer(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHTMLPageBuffer()
	}
}

func TestGetHTMLPageBufferNoGotFile(t *testing.T) {
	if err := GetHTMLPageBufferNoGotFile(); err != nil {
		t.Error(err)
	}
}

/* go test -bench=GetHTMLPageBuffer */
func BenchmarkGetHTMLPageBufferNoGotFile(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHTMLPageBufferNoGotFile()
	}
}
