// Package binedit handles a huge file named "data.bin" that does not fit in memory.
// It creates a new file where every 7th byte is skipped.
package binedit

import (
	"io"
	"os"
)

const (
	d  = "output"
	ff = "data.bin"
	ft = "datawo7.bin"
)

// binEdit returns true all operations completed until the last byte.
// It any operation failed, it panics.
func binEdit() bool {
	// File in running directory
	file, err := os.Open(ff)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	// File is overwritten when it exists.
	wfile, err := os.Create(ft)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	b1 := make([]byte, 1)
	for err != io.EOF { // Until the end of the file
		for i := 1; i < 7; i++ {
			_, err = file.Read(b1)
			if err != io.EOF {
				if err != nil {
					panic(err)
				}
				_, err = wfile.Write(b1)
				if err != nil {
					panic(err)
				}
			}
		}
		// Discarding the 7th byte
		_, err = file.Read(b1)
		if err != io.EOF {
			if err != nil {
				panic(err)
			}
		}
	}
	// Exit code is EOF and not 0
	return err == io.EOF
}
