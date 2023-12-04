package leetcode

// @Title        0367.Valid-Perfect-Square_test.go
// @Description
// @Create       XdpCs 2023-12-04 17:18
// @Update       XdpCs 2023-12-04 17:18

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestIsPerfectSquare(t *testing.T) {
	testCases := []test.Case{
		{
			Arg:      16,
			Expected: true,
		},
		{
			Arg:      14,
			Expected: false,
		},
	}
	fmt.Println("------------------------LeetCode Problem 0367------------------------")
	for _, testCase := range testCases {
		result := isPerfectSquare(testCase.Arg.(int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
