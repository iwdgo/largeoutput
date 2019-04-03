package largeoutput

import (
	"testing"
)

func TestBinEdit(t *testing.T) {

	OutputDir()

	if !binEdit() {
		t.Errorf("bin edition failed \n")
	}

	if err := FileCompare("datawant.bin", "datawo7.bin"); err != nil {
		t.Errorf("%v", err)
	}
}

/*
Benchmarking file operations should be examined
*/
func BenchmarkBinEdit(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		binEdit()
	}
}
