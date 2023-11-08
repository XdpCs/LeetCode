package leetcode

// @Title        2236.Root-Equals-Sum-of-Children.go
// @Description  2236.Root-Equals-Sum-of-Children solution
// @Create       XdpCs 2023-09-19 17:31
// @Update       XdpCs 2023-11-08 16:47

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%+v", *t)
}

func checkTree(root *TreeNode) bool {
	return root.Val == (root.Left.Val + root.Right.Val)
}
