// Package sumofconverts takes a list of strings and returns the sum of the items representing an integer.
// Several methods are benchmarked.
package sumofconverts

import (
	"strconv"
)

// SumOfConverts uses an ordinary for loop.
func SumOfConverts(listofstrings []string) (sum int64) {
	for _, a := range listofstrings {
		// 0 returns an int for some reason
		if i, err := strconv.ParseInt(a, 10, 64); err == nil {
			sum += i
		}
	}
	return
}

// SumOfConvertsR uses recursion passing the slice of strings.
func SumOfConvertsR(restofstrings []string) int64 {
	i, err := strconv.ParseInt(restofstrings[0], 10, 64)
	if len(restofstrings) == 1 {
		return i
	}
	if err == nil {
		return i + SumOfConvertsR(restofstrings[1:])
	}
	return SumOfConvertsR(restofstrings[1:])
}

// SumOfConvertsRPtr uses recursion passing a pointer to the slice of strings which is edited.
// Benchmarking shows that is consumes about a tenth of the resources.
func SumOfConvertsRPtr(restofstrings *[]string) int64 {
	i, err := strconv.ParseInt((*restofstrings)[0], 10, 64)
	if len(*restofstrings) == 1 {
		return i // zero if item is not an int
	}
	if err == nil {
		*restofstrings = (*restofstrings)[1:]
		return i + SumOfConvertsRPtr(restofstrings)
	}
	*restofstrings = (*restofstrings)[1:]
	return SumOfConvertsRPtr(restofstrings)
}
