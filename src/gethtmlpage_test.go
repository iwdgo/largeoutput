package largeoutput

import (
	"testing"
	//"bytes"
	"bytes"
)

/*

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
	OutputDir()
	pfileName := "pagegot.html"
	BufferToFile(pfileName, bytes.NewBuffer(getHTMLPage()))
	if err := FileCompare("pagewant.html", pfileName, t.Name()); err != nil {
		t.Error(err) // Otherwise, no error is detected
	}
}

/* go test -bench=GetHTMLPageBuffer */
func BenchmarkGetHTMLPageBuffer(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		getHTMLPage()
	}
}

/* go test -run=TestGetHTMLPageBuffer */
func TestGetHTMLPageBufferNoGotFile(t *testing.T) {
	OutputDir() // for want file
	if err := BufferCompare(bytes.NewBuffer(getHTMLPage()), "pagewant.html", t.Name()); err != nil {
		t.Error(err) // Otherwise, no error is detected
	}
}

/* go test -bench=GetHTMLPageBuffer */
func BenchmarkGetHTMLPageBufferNoGotFile(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		getHTMLPage()
	}
}
