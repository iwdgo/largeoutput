// Package modulo37 outputs sequentially the integers from 1 to 99 but on some conditions prints a string instead:
//  - when the integer is a multiple of 3 print “Open” instead of the number,
//  - when it is a multiple of 7 print “Source” instead of the number,
//  - when it is a multiple of both 3 and 7 print “OpenSource” instead of the number.
package modulo37

import (
	"fmt"
)

func modulo37() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Printf("Open")
		}
		if i%7 == 0 {
			fmt.Printf("Source")
		}
		if (i%3 != 0) && (i%7 != 0) {
			fmt.Printf("%d\n", i)
		} else {
			fmt.Printf("\n")
		}
	}
}
