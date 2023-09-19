package leetcode

// @Title        2236.Root-Equals-Sum-of-Children.go
// @Description  2236.Root-Equals-Sum-of-Children solution
// @Create       XdpCs 2023-09-19 17:31
// @Update       XdpCs 2023-09-19 17:31

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return root.Val == (root.Left.Val + root.Right.Val)
}
