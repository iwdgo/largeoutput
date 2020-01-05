package binedit

import (
	"github.com/iwdgo/testingfiles"
	"os"
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

func TestBinEditFileFailure(t *testing.T) {

	testingfiles.OutputDir("output")
	// Set read only
	err := os.Remove("datawo7.bin")
	if err != nil {
		t.Error(err)
	}
	fd, err := os.Create("datawo7.bin")
	if err != nil {
		t.Error(err)
	}
	err = fd.Close()
	if err != nil {
		t.Error(err)
	}
	err = os.Chmod("datawo7.bin", 400)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := recover(); err != nil {
			//!os.IsPermission(err.(error)) {
			t.Errorf("Recovering failed with %v", err)
		}
	}()

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
