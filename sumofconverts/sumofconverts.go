// Package sumofconverts takes a list of strings and returns the sum of the list items
// that represents an integer (skipping the other items)
// A recursive and non-recursive version are benchmarked.
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

func sumOfConverts() (sum int64) {
	for _, a := range listofstrings {
		// 0 returns an int for some reason
		if i, err := strconv.ParseInt(a, 10, 64); err == nil {
			sum += i
		}
	}
	return
}
