package leetcode

import (
	"fmt"
	"testing"
)

// @Title        1512.Number-of-Good-Pairs_test.go
// @Description  1512.Number-of-Good-Pairs_test solution test
// @Create       XdpCs 2023-09-25 16:39
// @Update       XdpCs 2023-09-25 16:39

func TestNumIdenticalPairs(t *testing.T) {
	testCases := []struct {
		nums []int
	}{
		{[]int{1, 2, 3, 1, 1, 3}},
		{[]int{1, 1, 1, 1}},
		{[]int{1, 2, 3}},
	}

	fmt.Println("------------------------LeetCode Problem 1512------------------------")
	for _, testCase := range testCases {
		fmt.Printf("Input: %v Output: %d\n", testCase.nums, numIdenticalPairs(testCase.nums))
	}
}
