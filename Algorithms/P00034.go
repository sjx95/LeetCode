package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 8, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{8, 8, 8, 8, 8, 8}, 8))
	fmt.Println(searchRange([]int{0, 8, 8, 8, 8, 8}, 8))
	fmt.Println(searchRange([]int{0, 0, 8, 8, 8, 40}, 8))
	fmt.Println(searchRange([]int{}, 8))
}

func searchRange(nums []int, target int) []int {
	var l, r int
	for ll, lr := 0, len(nums)-1; ll < lr; {
		l = (ll + lr) / 2
		if nums[l] >= target {
			lr = l
		} else {
			ll = l + 1
		}
		l = ll
	}
	for rl, rr := 0, len(nums)-1; rl < rr; {
		r = (rl + rr + 1) / 2
		if nums[r] > target {
			rr = r - 1
		} else {
			rl = r
		}
		r = rr
	}
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if (nums[l] == target) && (nums[r] == target) {
		return []int{l, r}
	} else {
		return []int{-1, -1}
	}
}