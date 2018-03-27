package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tn7 := TreeNode{7, nil, nil}
	tn2 := TreeNode{2, nil, nil}
	tn1 := TreeNode{1, nil, nil}
	tn11 := TreeNode{11, &tn7, &tn2}
	tn13 := TreeNode{13, nil, nil}
	tn42 := TreeNode{4, nil, &tn1}
	tn41 := TreeNode{4, &tn11, nil}
	tn8 := TreeNode{8, &tn13, &tn42}
	tn5 := TreeNode{5, &tn41, &tn8}

	tn_3 := TreeNode{-3, nil, nil}
	tn_2 := TreeNode{-2, nil, &tn_3}

	fmt.Println(hasPathSum(&tn5, 22))
	fmt.Println(hasPathSum(&tn_2, -2))
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return true
		} else {
			return false
		}
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
