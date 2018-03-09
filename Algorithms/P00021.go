package main

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	switch {
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	}

	var ans *ListNode
	x1, x2 := l1, l2
	if l2.Val < l1.Val {
		ans = l2
		x2 = l2.Next
	} else {
		ans = l1
		x1 = l1.Next
	}

	tail := ans
	for x1 != nil && x2 != nil {
		if x1.Val < x2.Val {
			tail.Next = x1
			x1 = x1.Next
		} else {
			tail.Next = x2
			x2 = x2.Next
		}
		tail = tail.Next
	}

	if x1 == nil {
		tail.Next = x2
	} else {
		tail.Next = x1
	}

	return ans
}

func main() {
	l1 := gen([]int{1, 2, 3, 5, 8})
	l2 := gen([]int{-1})
	ans := mergeTwoLists(l1, l2)
	for ans != nil {
		fmt.Print(ans.Val, "->")
		ans = ans.Next
	}
}

func gen(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var tmp []*ListNode
	for _, val := range arr {
		tln := ListNode{val, nil}
		tmp = append(tmp, &tln)
	}
	for i, val := range tmp {
		if i == 0 {
			continue
		}
		tmp[i-1].Next = val
	}
	return tmp[0]
}

type ListNode struct {
	Val  int
	Next *ListNode
}
