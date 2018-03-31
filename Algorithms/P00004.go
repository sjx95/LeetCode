package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1}, []int{2, 3}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
	}
	for _, n := range nums1[i:] {
		nums = append(nums, n)
	}
	for _, n := range nums2[j:] {
		nums = append(nums, n)
	}
	fmt.Println(nums)
	count := len(nums1) + len(nums2)
	if count%2 == 0 {
		return float64((nums[count/2 - 1] + nums[count/2])) / 2
	} else {
		return float64(nums[count/2])
	}
}
