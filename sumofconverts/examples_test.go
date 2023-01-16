package sumofconverts_test

import (
	"fmt"
	"github.com/iwdgo/largeoutput/sumofconverts"
)

var listofstrings = []string{
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it yes",
	"5",
	"Outlook good",
	"Yes",
	"Signs point to yes",
	"Reply hazy try again",
	"Ask again later",
	"7",
	"Cannot predict now",
	"Concentrate and ask again",
	"Don't count on it",
	"2",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

func ExampleSumOfConverts() {
	fmt.Println(sumofconverts.SumOfConverts(listofstrings))
	// Output: 14
}

func ExampleSumOfConvertsR() {
	fmt.Println(sumofconverts.SumOfConvertsR(listofstrings))
	// Output: 14
}

// Although the list must be first duplicated, this method consumes about a tenth of the resources.
func ExampleSumOfConvertsRPtr() {
	l := listofstrings // Duplicate as it is edited in the recursive algorithm
	fmt.Println(sumofconverts.SumOfConvertsRPtr(&l))
	// Output: 14
}
