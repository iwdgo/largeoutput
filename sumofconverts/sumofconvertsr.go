package sumofconverts

import (
	"strconv"
)

func stringstoint(restofstrings []string) int64 {
	i, err := strconv.ParseInt(restofstrings[0], 10, 64)
	if err == nil {
		return i + stringstoint(restofstrings[1:])
	}
	if len(restofstrings) == 1 {
		return 0
	}
	return stringstoint(restofstrings[1:])
}

// sumOfConvertsR is the recursive version of sumOfConverts.
func sumOfConvertsR() int64 {
	return stringstoint(listofstrings)
}
