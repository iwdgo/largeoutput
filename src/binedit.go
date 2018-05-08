package largeoutput

import (
	"io"
	"os"
)

/* You have a huge file named "data.bin" that does not fit in memory, code a program that deletes every 7th byte of it.
Truncate can be used to change its size.
*/

func binEdit() bool {

	//File in running directory
	file, err := os.Open("data.bin")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	wfile, err := os.Create("datawo7.bin")
	defer wfile.Close()
	if err != nil {
		panic(err)
	}

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
