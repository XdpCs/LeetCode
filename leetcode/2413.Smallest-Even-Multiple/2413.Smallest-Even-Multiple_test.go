package leetcode

// @Title        2413.Smallest-Even-Multiple_test.go
// @Description  2413.Smallest-Even-Multiple solution test
// @Create       XdpCs 2023-09-19 17:31
// @Update       XdpCs 2023-11-13 10:47

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestSmallestEvenMultiple(t *testing.T) {
	testCases := []test.Case{
		{
			Arg:      5,
			Expected: 10,
		},
		{
			Arg:      6,
			Expected: 6,
		},
	}

	fmt.Println("------------------------LeetCode Problem 2413------------------------")
	for _, testCase := range testCases {
		result := smallestEvenMultiple(testCase.Arg.(int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
