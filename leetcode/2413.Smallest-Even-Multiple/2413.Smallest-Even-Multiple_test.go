package leetcode

// @Title        2413.Smallest-Even-Multiple_test.go
// @Description  2413.Smallest-Even-Multiple solution test
// @Create       XdpCs 2023-09-19 17:31
// @Update       XdpCs 2023-09-19 17:31

import (
	"fmt"
	"testing"
)

func TestSmallestEvenMultiple(t *testing.T) {
	testCases := []struct {
		n int
	}{
		{5},
		{6},
	}

	fmt.Println("------------------------LeetCode Problem 2413------------------------")
	for _, testCase := range testCases {
		fmt.Printf("Input: %v Output: %v\n", testCase.n, smallestEvenMultiple(testCase.n))
	}
}
