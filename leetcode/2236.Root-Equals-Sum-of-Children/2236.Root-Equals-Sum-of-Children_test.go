package leetcode

// @Title        2236.Root-Equals-Sum-of-Children_test.go
// @Description  2236.Root-Equals-Sum-of-Children test
// @Create       XdpCs 2023-09-19 17:31
// @Update       XdpCs 2023-11-08 16:47

import (
	"fmt"
	"github.com/XdpCs/leetcode/util/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckTree(t *testing.T) {
	type arg struct {
		Root *TreeNode
	}
	testCases := []test.Case{
		{
			Arg: arg{
				Root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			Expect: true,
		},
		{
			Arg: arg{
				Root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			Expect: false,
		},
	}

	fmt.Println("------------------------LeetCode Problem 2236------------------------")
	for _, testCase := range testCases {
		result := checkTree(testCase.Arg.(arg).Root)
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}
