// Package binedit handles a huge file named "data.bin" that does not fit in memory.
// It creates a new file where every 7th byte is skipped.
package binedit

import (
	"fmt"
	"io"
	"os"
)

const (
	d  = "output"      // working directory
	ff = "data.bin"    // source file
	ft = "datawo7.bin" // edited file
	r  = 7             // order of byte to remove
)

// binEdit returns true when all operations completed until the last byte.
// If any operation but closing a file fails, it panics.
func binEdit() bool {
	// File in running directory
	file, err := os.Open(ff)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// File is overwritten when it exists.
	wfile, err := os.Create(ft)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = wfile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	b1 := make([]byte, 1)
	for err != io.EOF {
		for i := 1; i < r; i++ {
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
		// Discarding one byte
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
