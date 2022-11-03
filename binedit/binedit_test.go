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

	// Set working directory
	testingfiles.OutputDir(d)
	if !binEdit() {
		t.Errorf("bin edition failed\n")
	}
	if err := testingfiles.FileCompare(ft, "datawant.bin"); err != nil {
		t.Error(err)
	}
	fi, err := os.Stat(ff)
	if err != nil {
		t.Fatal(err)
	}
	fo, err := os.Stat(ft)
	if err != nil {
		t.Fatal(err)
	}
	if si := fi.Size() - (fi.Size() / r); si == fo.Size() {
		t.Logf("%d bytes removed from %d = %d", fi.Size()/7, fi.Size(), fo.Size())
	} else {
		t.Fatalf("output size: got %v, want %v", si, fo.Size())
	}
	err = os.RemoveAll(ft)
	if err != nil {
		t.Log(err)
	}
}

func TestBinEdit_readonly(t *testing.T) {
	testingfiles.OutputDir(d)
	// Discard error file could be absent
	_ = os.Remove(ft)
	fd, err := os.Create(ft)
	if err != nil {
		t.Fatal(err)
	}
	err = fd.Close()
	if err != nil {
		t.Error(err)
	}
	// Set output file read only
	var p os.FileMode = 0400
	err = os.Chmod(ft, p)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
		err = os.RemoveAll(ft)
		if err != nil {
			t.Log(err)
			t.Logf("%s might be read-only. Remove manually.", ft)
		}
	}()

	if !binEdit() {
		// Because of panic-ing, this code must be unreachable
		t.Errorf("bin didn't panic")
	}

	t.Logf("No panic and no error for perm %v", p)
	if err := testingfiles.FileCompare(ft, "datawant.bin"); err != nil {
		t.Error(err)
	}
}

func BenchmarkBinEdit(b *testing.B) {
	b.Skip("permission denied")
	for n := 0; n < b.N; n++ {
		if !binEdit() {
			b.Fatalf("bin edit failed on %d run", n)
		}
	}
}
