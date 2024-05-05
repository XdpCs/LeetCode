package main

// @Title        0326.Power-Of-Three_test.go
// @Description  0326.Power-Of-Three test
// @Create       XdpCs 2023-11-13 11:04
// @Update       XdpCs 2023-11-13 11:04

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

var testCases = []test.Case{
	{
		Arg:      27,
		Expected: true,
	},
	{
		Arg:      0,
		Expected: false,
	},
	{
		Arg:      9,
		Expected: true,
	},
	{
		Arg:      45,
		Expected: false,
	},
}

func TestGetMaxInt32PowerOfThree(t *testing.T) {
	result := GetMaxInt32PowerOfThree()
	t.Log(result)
}

func TestIsPowerOfThree(t *testing.T) {
	fmt.Println("------------------------LeetCode Problem 0326------------------------")
	for _, testCase := range testCases {
		result := isPowerOfThree(testCase.Arg.(int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}

func TestIsPowerOfThreeMath(t *testing.T) {
	fmt.Println("------------------------LeetCode Problem 0326------------------------")
	for _, testCase := range testCases {
		result := isPowerOfThreeMath(testCase.Arg.(int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
func TestIsPowerOfThreeLoop(t *testing.T) {
	fmt.Println("------------------------LeetCode Problem 0326------------------------")
	for _, testCase := range testCases {
		result := isPowerOfThreeLoop(testCase.Arg.(int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
