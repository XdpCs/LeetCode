package leetcode

// @Title        2413.Smallest-Even-Multiple_test.go
// @Description  2413.Smallest-Even-Multiple solution test
// @Create       XdpCs 2023-09-19 17:31
// @Update       XdpCs 2023-11-08 16:47

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestSmallestEvenMultiple(t *testing.T) {
	type arg struct {
		n int
	}
	testCases := []test.Case{
		{
			Arg:    arg{5},
			Expect: 10,
		},
		{
			Arg:    arg{6},
			Expect: 6,
		},
	}

	fmt.Println("------------------------LeetCode Problem 2413------------------------")
	for _, testCase := range testCases {
		result := smallestEvenMultiple(testCase.Arg.(arg).n)
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}
