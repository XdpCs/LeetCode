package leetcode

import (
	"fmt"
	"testing"
)

// @Title        1534.Count-Good-Triplets_test.go
// @Description
// @Create       XdpCs 2023-09-25 19:12
// @Update       XdpCs 2023-09-25 19:12

func TestCountGoodTriplets(t *testing.T) {
	testCases := []struct {
		arr []int
		a   int
		b   int
		c   int
	}{
		{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3},
		{[]int{1, 1, 2, 2, 3}, 0, 0, 1},
	}

	fmt.Println("------------------------LeetCode Problem 1534------------------------")
	for _, testCase := range testCases {
		fmt.Printf("Input: %v Output: %d\n", testCase.arr, countGoodTriplets(testCase.arr, testCase.a, testCase.b, testCase.c))
	}
}
