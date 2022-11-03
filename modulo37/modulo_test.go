package modulo37

import (
	"bytes"
	"fmt"
	"github.com/iwdgo/testingfiles"
	"io"
	"os"
	"testing"
	"time"
)

/*

About testing

Example would be like Test...
Log format, if used, and includes date, time...
pfile := iotest.NewWriteLogger(t.Name(),os.Stdout) and no valid reference can be created

Go test // Output: is awkward considering the size of the output.

To handle output, you can write the func with a io.Writer parameter like below, but it requires to update or to write
code with test in mind.

func modulo37(f io.Writer) {
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Fprint(f,"Open")
		}
		if i%7 == 0 {
			fmt.Fprint(f,"Source")
		}
		if (i%3 != 0) && (i%7 != 0) {
			fmt.Fprintf(f,"%d\n",i)
		} else {
			fmt.Fprintln(f)
		}
	}
}

Without access to the source code, the output can be piped to a file.
A reference file is used for comparison purposes.

*/

// Test is piping output to a file which is checked against reference.
func TestModulo37(t *testing.T) {
	testingfiles.OutputDir("output")
	prodFileName := "moduloprod.txt"
	pfile, err := os.Create(prodFileName)
	if err != nil {
		t.Errorf("%v", err)
	}

	// Capture stdout.
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "modulo : piping failed with ", err)
		os.Exit(1)
	}
	os.Stdout = w
	outC := make(chan []byte)
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		_ = r.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "testing: copying pipe: %v\n", err)
			os.Exit(1)
		}
		outC <- buf.Bytes() //.String()
	}()

	start := time.Now()
	ok := true

	// Clean up in a deferred call so we can recover if the example panics.
	defer func() {
		err := recover()
		if err != nil { // If here because of panic
			t.Error(err)
			panic(err) // Testing output has no value
		}

		dstr := fmt.Sprintf("%.4fs", time.Since(start).Seconds())

		// Close pipe, restore stdout, get output.
		_ = w.Close()
		os.Stdout = stdout // Restoring Stdout
		out := <-outC
		_, _ = pfile.Write(out)
		_ = pfile.Close()

		if err = testingfiles.FileCompare(prodFileName, "modulowant.txt"); err != nil {
			t.Errorf("%s : %v\n", dstr, err)
		}

		ok = err == nil
	}()

	if !ok {
		t.Errorf("Opening pipe failed with %v", err)
	}
	modulo37()
	// All output handling is in defer
}

func BenchmarkModulo37(b *testing.B) {
	prodFileName := "modulobench.txt"
	pfile, err := os.Create(prodFileName)
	if err != nil {
		b.Errorf("%v", err)
	}

	// Capture stdout.
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "modulo : piping failed with ", err)
		os.Exit(1)
	}
	os.Stdout = w
	outC := make(chan []byte)
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		_ = r.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "testing: copying pipe: %v\n", err)
			os.Exit(1)
		}
		outC <- buf.Bytes() //.String()
	}()

	/* Clean up in a deferred call so we can recover if the example panics. */
	defer func() {
		err := recover()
		if err != nil { // If here because of panic
			b.Error(err)
			panic(err) // Testing output has no value
		}

		// Close pipe, restore stdout, get output.
		_ = w.Close()
		os.Stdout = stdout // Restoring Stdout
		out := <-outC
		_, _ = pfile.Write(out)
		_ = pfile.Close()
	}()

	// run the function b.N times
	for n := 0; n < b.N; n++ {
		modulo37()
	}
}
