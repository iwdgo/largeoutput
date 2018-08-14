package largeoutput

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

/*

Example would be like Test...
Log format, if used, and includes date, time...
pfile := iotest.NewWriteLogger(t.Name(),os.Stdout) and no valid reference can be created

You can conceive your test to pass the produced file and all output is then fmt.Fprintf(pfile,...)
func TestModulo37(t *testing.T) {
	prodFileName := "moduloprod.txt"
	pfile, err := os.Create(prodFileName)
	defer pfile.Close()
	if err != nil {
		t.Errorf("Produced file creation %s failed with %v",prodFileName,err)
	}


	modulo37(pfile) // t.Log is using stdErr and looks confusing
	pfile.Close()

	FileCompare(t,"moduloref.txt",prodFileName,"modulo print")
}
*/

func TestModulo37(t *testing.T) {
	OutputDir()
	prodFileName := "moduloprod.txt"
	pfile, err := os.Create(prodFileName)
	if err != nil {
		t.Errorf("produced file creation %s failed with %v", prodFileName, err)
	}

	// Capture stdout.
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "modulo : piping failed with ", err)
		os.Exit(1)
	}
	os.Stdout = w
	outC := make(chan []byte)
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		r.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "testing: copying pipe: %v\n", err)
			os.Exit(1)
		}
		outC <- buf.Bytes() //.String()
	}()

	start := time.Now()
	ok := true

	/* Clean up in a deferred call so we can recover if the example panics. */
	defer func() {
		err := recover()
		if err != nil { // If here because of panic
			t.Error(err)
			panic(err) // Testing output has no value
		}

		dstr := fmt.Sprintf("%.4fs", time.Since(start).Seconds())

		// Close pipe, restore stdout, get output.
		w.Close()
		os.Stdout = stdout // Restoring Stdout
		out := <-outC
		pfile.Write(out)
		pfile.Close()

		fail := FileCompare("modulowant.txt", prodFileName, "modulo print")
		if fail != nil {
			t.Errorf("%s : %v\n", dstr, fail)
		}

		ok = fail == nil && err == nil
	}()

	ok = ok // otherwise -vet=off is needed. Cf. https://github.com/golang/go/issues/25720
	if err != nil {
		t.Errorf("Opening pipe failed with %v", err)
	}
	modulo37()
	// All output handling is in defer

}

/* */
func BenchmarkModulo37(b *testing.B) {
	prodFileName := "modulobench.txt"
	pfile, err := os.Create(prodFileName)
	if err != nil {
		b.Errorf("produced file creation %s failed with %v", prodFileName, err)
	}

	// Capture stdout.
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "modulo : piping failed with ", err)
		os.Exit(1)
	}
	os.Stdout = w
	outC := make(chan []byte)
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		r.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "testing: copying pipe: %v\n", err)
			os.Exit(1)
		}
		outC <- buf.Bytes() //.String()
	}()

	ok := true
	/* Clean up in a deferred call so we can recover if the example panics. */
	defer func() {
		err := recover()
		if err != nil { // If here because of panic
			b.Error(err)
			panic(err) // Testing output has no value
		}

		// Close pipe, restore stdout, get output.
		w.Close()
		os.Stdout = stdout // Restoring Stdout
		out := <-outC
		pfile.Write(out)
		pfile.Close()

		/* File content is not tested during benchmark */
		ok = err == nil
	}()
	/* */
	ok = ok // otherwise -vet=off is needed. Cf. https://github.com/golang/go/issues/25720
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		modulo37()
	}
}
