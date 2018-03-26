package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestValidParentheses(")(()))"))
}

func longestValidParentheses(s string) int {
	maxlen := 0
	countL, l := 0, 0
	for _, ch := range s {
		if ch == '(' {
			countL++
		}
		if ch == ')' {
			if countL > 0 {
				countL--
			} else {
				countL = 0
				l = 0
				continue
			}
		}
		l++
		if countL == 0 && l > maxlen {
			maxlen = l
		}
	}
	countR, l := 0, 0
	for i := len(s); i != 0; {
		i--
		if s[i] == ')' {
			countR++
		}
		if s[i] == '(' {
			if countR > 0 {
				countR--
			} else {
				countR = 0
				l = 0
				continue
			}
		}
		l++
		if countR == 0 && l > maxlen {
			maxlen = l
		}

	}

	return maxlen
}
