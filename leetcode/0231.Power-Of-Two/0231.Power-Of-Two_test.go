package leetcode

// @Title        0231.Power-Of-Two_test.go
// @Description  0231.Power-Of-Two test
// @Create       XdpCs 2023-11-13 11:16
// @Update       XdpCs 2023-11-13 11:16

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

var testCases = []test.Case{
	{
		Arg:    1,
		Expect: true,
	},
	{
		Arg:    16,
		Expect: true,
	},
	{
		Arg:    3,
		Expect: false,
	},
	{
		Arg:    4,
		Expect: true,
	},
	{
		Arg:    5,
		Expect: false,
	},
}

func TestIsPowerOfTwo(t *testing.T) {
	fmt.Println("------------------------LeetCode Problem 0231------------------------")
	for _, testCase := range testCases {
		result := isPowerOfTwo(testCase.Arg.(int))
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}

func TestIsPowerOfTwoStorage(t *testing.T) {
	fmt.Println("------------------------LeetCode Problem 0231------------------------")
	for _, testCase := range testCases {
		result := isPowerOfTwoStorage(testCase.Arg.(int))
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}

func TestIsPowerOfTwoMaxNum(t *testing.T) {
	fmt.Println("------------------------LeetCode Problem 0231------------------------")
	for _, testCase := range testCases {
		result := isPowerOfTwoMaxNum(testCase.Arg.(int))
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}
