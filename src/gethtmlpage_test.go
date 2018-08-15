package largeoutput

import (
	"testing"
	//"bytes"
	"bytes"
	"os"
)

/* Buffer beats String easily
goos: windows
goarch: amd64
BenchmarkGetHTMLPageString-4          10         322947190 ns/op
BenchmarkGetHTMLPageBuffer-4          10         186943420 ns/op
*/
func CreateFileFromString(fname string, content []byte) {
	wfile, err := os.Create(fname)
	defer wfile.Close()
	if err != nil {
		panic(err)
	}

	_, err = wfile.Write(content)
	if err != nil {
		panic(err)
	}

	//fmt.Println(getHTMLPage())
}

func BufferToFile(fname string, content *bytes.Buffer) {
	wfile, err := os.Create(fname)
	defer wfile.Close()
	if err != nil {
		panic(err)
	}

	_, err = wfile.Write(content.Bytes())
	if err != nil {
		panic(err)
	}
}

/* go test -run=TestGetHTMLPageString */
func TestGetHTMLPageString(t *testing.T) {
	OutputDir()
	pfileName := "pagegot.html"
	CreateFileFromString(pfileName, getHTMLPage())
	if err := FileCompare("pagewant.html", pfileName, t.Name()); err != nil {
		t.Error(err) // Otherwise, no error is detected
	}
}

/* go test -bench=GetHTMLPageString */
func BenchmarkGetHTMLPageString(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		getHTMLPage()
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
