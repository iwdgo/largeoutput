package binedit

import (
	"github.com/iwdgo/testingfiles"
	"os"
	"testing"
)

func TestBinEdit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()

	// Sets working directory
	testingfiles.OutputDir(d)
	if !binEdit() {
		t.Errorf("bin edition failed\n")
	}
	if err := testingfiles.FileCompare(ft, "datawant.bin"); err != nil {
		t.Errorf("%v", err)
	}
	_ = os.RemoveAll(ft)
}

func TestBinEdit_readonly(t *testing.T) {
	testingfiles.OutputDir(d)
	// Set read only
	err := os.Remove(ft)
	if err != nil {
		t.Error(err)
	}
	fd, err := os.Create(ft)
	if err != nil {
		t.Error(err)
	}
	err = fd.Close()
	if err != nil {
		t.Error(err)
	}
	var p os.FileMode = 0400
	err = os.Chmod(ft, p)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()

	if !binEdit() {
		// Because of panic-ing, this code must be unreachable
		t.Errorf("bin didn't panic")
	}

	t.Logf("No panic and no error for perm %v", p)
	if err := testingfiles.FileCompare(ft, "datawant.bin"); err != nil {
		t.Errorf("%v", err)
	}
}

func BenchmarkBinEdit(b *testing.B) {
	b.Skip("permission denied")
	// Run the function b.N times
	for n := 0; n < b.N; n++ {
		if !binEdit() {
			b.Fatalf("bin edit failed on %d run", n)
		}
	}
}
