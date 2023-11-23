package leetcode

// @Title        0867.Transpose-Matrix_test.go
// @Description
// @Create       XdpCs 2023-11-22 20:57
// @Update       XdpCs 2023-11-22 20:57

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestTranspose(t *testing.T) {
	testCase := []test.Case{
		{
			Arg:      [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			Expected: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			Arg:      [][]int{{1, 2, 3}, {4, 5, 6}},
			Expected: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
	}
	fmt.Println("------------------------LeetCode Problem 0867------------------------")
	for _, testCase := range testCase {
		result := transpose(testCase.Arg.([][]int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
