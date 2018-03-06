package main

import "fmt"

func main() {
	fmt.Println(myAtoi(" b2147483649"))
}

func myAtoi(str string) int {
	for len(str) != 0 && str[0] == ' ' {
		str = str[1:]
	}
	if len(str) == 0 || !(str[0] == '+' || str[0] == '-' || (str[0] >= '0' && str[0] <= '9')) {
		return 0;
	}

	var ans int
	unit := 1
	switch str[0] {
	case '-':
		unit = -1
		str = str[1:]
	case '+':
		str = str[1:]
	}
	for i, ch := range str {
		if !(ch >= '0' && ch <= '9') {
			str = str[:i]
			break;
		}
	}

	for _, ch := range str {
		x := int(ch - '0')

		switch {
		case (-214748364 < ans) && (ans < 214748364):
			fallthrough
		case ans == -214748364 && x <= 8:
			fallthrough
		case ans == 214748364 && x <= 7:
			ans = ans * 10 + unit * x
		case unit == 1:
			return 2147483647
		case unit == -1:
			return -2147483648
		}
	}
	return ans;
}
