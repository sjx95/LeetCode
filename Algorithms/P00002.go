package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	b := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	ans := addTwoNumbers(a, b)
	if ans != nil {
		fmt.Print(ans.Val)
	}
	for ans.Next != nil {
		ans = ans.Next;
		fmt.Print(" -> ", ans.Val)
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	tail := head
	val := 0

	for (l1 != nil) && (l2 != nil) {
		val += l1.Val + l2.Val
		if head == nil {
			tail = &ListNode{val % 10, nil}
			head = tail
		} else {
			tail.Next = &ListNode{val % 10, nil}
			tail = tail.Next
		}
		val /= 10
		l1, l2 = l1.Next, l2.Next
	}

	for l1 != nil {
		val += l1.Val
		if head == nil {
			tail = &ListNode{val % 10, nil}
			head = tail
		} else {
			tail.Next = &ListNode{val % 10, nil}
			tail = tail.Next
		}
		val /= 10
		l1 = l1.Next
	}

	for l2 != nil {
		val += l2.Val
		if head == nil {
			tail = &ListNode{val % 10, nil}
			head = tail
		} else {
			tail.Next = &ListNode{val % 10, nil}
			tail = tail.Next
		}
		val /= 10
		l2 = l2.Next
	}

	if val > 0 {
		tail.Next = &ListNode{val, nil}
	}

	return head
}
