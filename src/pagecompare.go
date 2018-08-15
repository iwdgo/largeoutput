package largeoutput

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func FileCompare(want, got, caller string) error {
	rfile, err := os.Open(want)
	defer rfile.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Reference file %s open failed with %v", want, err))
	}

	pfile, err := os.Open(got)
	defer pfile.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Profuced file %s open failed with %v", got, err))
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

		if !bytes.Equal(b1, b2) {
			return errors.New(fmt.Sprintf(caller+" : got %q, want %q at %d", b1, b2, index))
			break
		}
		index++
	}
	// EOF on reference file has been reached, let us check the produced file
	_, err = pfile.Read(b2)
	if err != io.EOF { // If EOF produced file is too short
		rfileInfo, _ := rfile.Stat()
		return errors.New(fmt.Sprintf(caller+" : produced file is too short by %d", rfileInfo.Size()-int64(index)))
	}
	return nil
}

// Buffer is compared to file. If an error occurs, got file is created, otherwise nil is returned.
func BufferCompare(got *bytes.Buffer, want, caller string) error {
	rfile, err := os.Open(want)
	defer rfile.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Reference file %s open failed with %v", want, err))
	}

	b1, b2 := make([]byte, 1), make([]byte, 1)
	index := 0          // Index in file to locate error
	for err != io.EOF { // Until the end of the file
		_, err = rfile.Read(b1)
		if err != io.EOF { // While not EOF, read the other file too
			if err != nil {
				return err
			}

			_, err = got.Read(b2)
			if err != nil { // If EOF produced file is too short
				return err
			}
		}

		if !bytes.Equal(b1, b2) {
			BufferToFile("got"+caller+".html", got)
			return errors.New(fmt.Sprintf(caller+" : got %q, want %q at %d", b1, b2, index))
			break
		}
		index++
	}
	// EOF on reference file has been reached, let us check the produced file
	_, err = got.Read(b2)
	if err != io.EOF { // If EOF produced file is too short
		BufferToFile("got"+caller+".html", got)
		rfileInfo, _ := rfile.Stat()
		return errors.New(fmt.Sprintf(caller+" : produced file is too short by %d", rfileInfo.Size()-int64(index)))
	}
	return nil
}
