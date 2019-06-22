// Write a function that takes a list of strings and returns the sum of the list items
// that represents an integer (skipping the other items)

package largeoutput

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

func sumOfConverts() int64 {
	var sum int64
	sum = 0
	for _, a := range listofstrings {
		i, err := strconv.ParseInt(a, 10, 64) // 0 returns an int for some reason
		if err == nil {
			sum = sum + i
		}
	}
	return sum
}
