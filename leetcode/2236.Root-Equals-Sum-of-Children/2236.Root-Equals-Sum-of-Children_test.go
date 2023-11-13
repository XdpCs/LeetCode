package leetcode

// @Title        2236.Root-Equals-Sum-of-Children_test.go
// @Description  2236.Root-Equals-Sum-of-Children test
// @Create       XdpCs 2023-09-19 17:31
// @Update       XdpCs 2023-11-13 10:21

import (
	"fmt"
	"testing"

	"github.com/XdpCs/leetcode/util/test"

	"github.com/stretchr/testify/assert"
)

func TestCheckTree(t *testing.T) {
	testCases := []test.Case{
		{
			Arg: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},

			Expect: true,
		},
		{
			Arg: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Expect: false,
		},
	}

	fmt.Println("------------------------LeetCode Problem 2236------------------------")
	for _, testCase := range testCases {
		result := checkTree(testCase.Arg.(*TreeNode))
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}
