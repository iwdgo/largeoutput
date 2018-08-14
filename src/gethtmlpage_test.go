package largeoutput

import (
	//"fmt"
	//"io/ioutil"
	"os"
	"testing"
	//"bytes"
)

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

/* go test -run=TestGetHTMLPage */
func TestGetHTMLPage(t *testing.T) {
	OutputDir()
	pfileName := "pagegot.html"
	s := getHTMLPage()
	CreateFileFromString(pfileName, s)
	if err := FileCompare("pagewant.html", pfileName, "change html page"); err != nil {
		t.Error(err) // Otherwise, no error is detected
	}
}

/* */
func BenchmarkGetHTMLPage(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		getHTMLPage()
	}
}
