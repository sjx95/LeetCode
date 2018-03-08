package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return math.MaxInt32
	}

	sign := 1
	var pdividend uint
	var pdivisor uint
	if dividend < 0 {
		pdividend = uint(-dividend)
		sign *= -1
	} else {
		pdividend = uint(dividend)
	}
	if divisor < 0 {
		pdivisor = uint(-divisor)
		sign *= -1
	} else {
		pdivisor = uint(divisor)
	}

	ans := 0
	var sum uint = 0
	for sum + pdivisor <= pdividend{
		i, psum := 1, pdivisor
		for psum + psum + sum <= pdividend {
			psum += psum
			i += i
		}
		ans += i
		sum += psum
	}

	ans *= sign

	return ans
}

func main() {
	fmt.Println(divide(5, 3))
}
