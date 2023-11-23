package leetcode

// @Title        1470.Shuffle-The-Array_test.go
// @Description	 1470.Shuffle-The-Array test
// @Create       XdpCs 2023-11-22 20:37
// @Update       XdpCs 2023-11-22 20:37

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestShuffle(t *testing.T) {
	type arg struct {
		nums []int
		n    int
	}
	testCases := []test.Case{
		{
			Arg: arg{
				nums: []int{2, 5, 1, 3, 4, 7},
				n:    3,
			},
			Expected: []int{2, 3, 5, 4, 1, 7},
		},
		{
			Arg: arg{
				nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
				n:    4,
			},
			Expected: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			Arg: arg{
				nums: []int{1, 1, 2, 2},
				n:    2,
			},
			Expected: []int{1, 2, 1, 2},
		},
	}
	fmt.Println("------------------------LeetCode Problem 1470------------------------")
	for _, testCase := range testCases {
		result := shuffle(testCase.Arg.(arg).nums, testCase.Arg.(arg).n)
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
