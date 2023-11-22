package leetcode

// @Title        1512.Number-of-Good-Pairs_test.go
// @Description  1512.Number-of-Good-Pairs_test solution test
// @Create       XdpCs 2023-09-25 16:39
// @Update       XdpCs 2023-11-13 10:23

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestNumIdenticalPairs(t *testing.T) {
	testCases := []test.Case{
		{
			Arg:      []int{1, 2, 3, 1, 1, 3},
			Expected: 4,
		},
		{
			Arg:      []int{1, 1, 1, 1},
			Expected: 6,
		},
		{
			Arg:      []int{1, 2, 3},
			Expected: 0,
		},
	}

	fmt.Println("------------------------LeetCode Problem 1512------------------------")
	for _, testCase := range testCases {
		result := numIdenticalPairs(testCase.Arg.([]int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
