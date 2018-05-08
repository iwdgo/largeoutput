package largeoutput

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/* All data files are stored outside the code in a "output" subdirectory */
func OutputDir() {
	ex, err := os.Getwd() // Executable() is where Go runs not where the files are created
	if err != nil {
		panic(err)
	}
	// ex := os.TempDir() // like the name says is a place for temporary files
	if filepath.Base(ex) != "output" { // No need to change
		err = os.Chdir(ex + "/output")
		if err != nil {
			panic(err) // output subdirectory is probably missing
		}
	}
}

/*
 To check large outputs of a test, a file storage is more convenient or required.
 This routine compares produced output to a reference file and returns an error.
 Stdout must be piped to the produced file.
*/

func FileCompare(ref, prod, testName string) error {
	rfile, err := os.Open(ref)
	defer rfile.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Reference file %s open failed with %v", ref, err))
	}

	pfile, err := os.Open(prod)
	defer pfile.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Profuced file %s open failed with %v", prod, err))
	}

	b1, b2 := make([]byte, 1), make([]byte, 1)
	index := 0          // Index in file to locate error
	for err != io.EOF { // Until the end of the file
		_, err = rfile.Read(b1)
		if err != io.EOF { // While not EOF, read the other file too
			if err != nil {
				return err
			}

			_, err = pfile.Read(b2)
			if err != nil { // If EOF produced file is too short
				return err
			}
		}

		// Discarding the 7th byte
		if !bytes.Equal(b1, b2) {
			return errors.New(fmt.Sprintf(testName+" : got %q, want %q at %d", b1, b2, index))
			break
		}
		index++
	}
	// EOF on reference file has been reached, let us check the produced file
	_, err = pfile.Read(b2)
	if err != io.EOF { // If EOF produced file is too short
		rfileInfo, _ := rfile.Stat()
		return errors.New(fmt.Sprintf(testName+" : produced file is too short by %d", rfileInfo.Size()-int64(index)))
	}
	return nil
}
