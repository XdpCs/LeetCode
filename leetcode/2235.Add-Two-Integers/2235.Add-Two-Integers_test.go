package leetcode

import (
	"fmt"
	"testing"
)

// @Title        2235.Add-Two-Integers_test.go
// @Description  2235.Add-Two-Integers solution test
// @Create       XdpCs 2023-09-18 21:05
// @Update       XdpCs 2023-09-18 21:05

func TestSum(t *testing.T) {
	testCases := []struct {
		num1 int
		num2 int
	}{
		{1118, 1114},
		{12, 5},
		{-10, 4},
		{0, 0},
	}

	fmt.Println("------------------------LeetCode Problem 2235------------------------")
	for _, testCase := range testCases {
		fmt.Printf("Input: %d, %d Output: %d\n", testCase.num1, testCase.num2, sum(testCase.num1, testCase.num2))
	}
}
