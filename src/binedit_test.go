package largeoutput

import (
	"github.com/iwdgo/testingfiles"
	"testing"
)

func TestBinEdit(t *testing.T) {

	testingfiles.OutputDir("output")

	if !binEdit() {
		t.Errorf("bin edition failed \n")
	}

	if err := testingfiles.FileCompare("datawant.bin", "datawo7.bin"); err != nil {
		t.Errorf("%v", err)
	}
}

func BenchmarkBinEdit(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		if !binEdit() {
			b.Fatalf("bin edit failed on %d run", n)
		}
	}
}
