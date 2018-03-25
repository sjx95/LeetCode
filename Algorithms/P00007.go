package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(8463847412))
}

func reverse(x int) int {
	y := 0

	for x != 0 {
		switch {
		case y == 214748364 && x % 10 > 7:
			fallthrough
		case y > 214748364:
			fallthrough
		case y < -214748364:
			fallthrough
		case y == -214748364 && x % 10 < -8:
			return 0
		}

		y *= 10
		y += x % 10
		x /= 10
	}
	return y
}