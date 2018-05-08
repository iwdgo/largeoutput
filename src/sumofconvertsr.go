package largeoutput

import (
	"strconv"
)

/*
Write a recursive version of the previous function (or an iterative version if you already did a recursive version).
*/
/*
func listofstringsR() []string {
	return []string{
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
}
*/
func stringstoint(restofstrings []string) int64 {
	i, err := strconv.ParseInt(restofstrings[0], 10, 64)
	if err == nil {
		return i + stringstoint(restofstrings[1:])
	} else if len(restofstrings) == 1 {
		return 0
	} else {
		return stringstoint(restofstrings[1:])
	}
}

func sumOfConvertsR() int64 {
	return stringstoint(listofstrings())
}
