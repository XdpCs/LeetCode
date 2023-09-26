package leetcode

import (
	"fmt"
	"testing"
)

// @Title        1486.XOR-Operation-in-an-Array_test.go
// @Description  1486.XOR-Operation-in-an-Array_test solution test
// @Create       XdpCs 2023-09-25 16:17
// @Update       XdpCs 2023-09-25 16:17

func TestXorOperation(t *testing.T) {
	testCases := []struct {
		n     int
		start int
	}{
		{5, 0},
		{4, 3},
		{1, 7},
		{10, 5},
	}

	fmt.Println("------------------------LeetCode Problem 1486------------------------")
	for _, testCase := range testCases {
		fmt.Printf("Input: %d, %d Output: %d\n", testCase.n, testCase.start, xorOperation(testCase.n, testCase.start))
	}
}
