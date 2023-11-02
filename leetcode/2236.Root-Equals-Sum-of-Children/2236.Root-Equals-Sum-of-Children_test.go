package leetcode

// @Title        2236.Root-Equals-Sum-of-Children_test.go
// @Description  2236.Root-Equals-Sum-of-Children test
// @Create       XdpCs 2023-09-19 17:31
// @Update       XdpCs 2023-09-19 17:31

import (
	"fmt"
	"testing"
)

func TestCheckTree(t *testing.T) {
	testCases := []struct {
		root *TreeNode
	}{
		{&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
			},
		}},
		{&TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		}},
	}

	fmt.Println("------------------------LeetCode Problem 2236------------------------")
	for _, testCase := range testCases {
		fmt.Printf("Input: %+v left: %+v right: %+v Output: %v\n", testCase.root.Val, testCase.root.Left.Val, testCase.root.Right.Val, checkTree(testCase.root))
	}
}
