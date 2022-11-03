/*
	Package sumofconverts takes a list of strings and returns the sum of the items representing an integer.

Recursive and non-recursive methods are implemented without significant changes in performance.
*/
package sumofconverts

import (
	"strconv"
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

const total = 14

// SumOfConverts uses an ordinary for loop
func SumOfConverts(listofstrings []string) (sum int64) {
	for _, a := range listofstrings {
		// 0 returns an int for some reason
		if i, err := strconv.ParseInt(a, 10, 64); err == nil {
			sum += i
		}
	}
	return
}

// SumOfConvertsR uses recursion passing the slice of strings
func SumOfConvertsR(restofstrings []string) int64 {
	i, err := strconv.ParseInt(restofstrings[0], 10, 64)
	if err == nil {
		return i + SumOfConvertsR(restofstrings[1:])
	}
	if len(restofstrings) == 1 {
		return 0
	}
	return SumOfConvertsR(restofstrings[1:])
}

// SumOfConvertsRPtr uses recursion passing a pointer to the slice of strings
func SumOfConvertsRPtr(restofstrings *[]string) int64 {
	i, err := strconv.ParseInt((*restofstrings)[0], 10, 64)
	if err == nil {
		*restofstrings = (*restofstrings)[1:]
		return i + SumOfConvertsRPtr(restofstrings)
	}
	if len(*restofstrings) == 1 {
		return 0
	}
	*restofstrings = (*restofstrings)[1:]
	return SumOfConvertsRPtr(restofstrings)
}
