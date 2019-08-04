package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,0,-1,0,-2,2}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	res := map[int]int{}
	for _, x := range(nums) {
		res[x] = res[x] + 1
	}
	xs := make([]int, 0, len(nums))
	for x := range(res) {
		xs = append(xs, x)
	}
	cs := make([]int, 0, len(xs))
	sort.Ints(xs)
	for _, x := range(xs) {
		cs = append(cs, res[x])
	}

	ans := make([][]int, 0, 128)
	for a := 0; a < len(xs); a++ {
		cs[a]--
		b := a
		if cs[a] == 0 {
			b++
		}
		for ; b < len(xs); b++ {
			cs[b]--
			c := b
			if cs[b] == 0 {
				c++
			}
			for ; c < len(xs); c++ {
				cs[c]--
				d := c
				if cs[c] == 0 {
					d++
				}
				for ; d < len(xs); d++ {
					data := []int{xs[a], xs[b], xs[c], xs[d]}
					sum := 0
					for _, x := range(data) {
						sum += x
					}
					if sum == target {
						ans = append(ans, data)
					}
				}
				cs[c]++
			}
			cs[b]++
		}
		cs[a]++
	}
	return ans
}


