package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

// @Title        0852.Peak-Index-In-A-Mountain-Array_test.go
// @Description  0852.Peak-Index-In-A-Mountain-Array test
// @Create       XdpCs 2023-11-23 12:08
// @Update       XdpCs 2023-11-23 12:08

func TestPeakIndexInMountainArray(t *testing.T) {
	testCases := []test.Case{
		{
			Arg:      []int{0, 1, 0},
			Expected: 1,
		},
		{
			Arg:      []int{0, 2, 1, 0},
			Expected: 1,
		},
		{
			Arg:      []int{0, 10, 5, 2},
			Expected: 1,
		},
	}
	fmt.Println("------------------------LeetCode Problem 0852------------------------")
	for _, testCase := range testCases {
		result := peakIndexInMountainArray(testCase.Arg.([]int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
