package leetcode

// @Title        1281.Subtract-the-Product-and-Sum-of-Digits-of-an-Integer_test.go
// @Description  1281.Subtract-the-Product-and-Sum-of-Digits-of-an-Integer test
// @Create       XdpCs 2023-11-06 12:23
// @Update       XdpCs 2023-11-06 12:23

import (
	"fmt"
	"testing"
)

func TestAddDigits(t *testing.T) {
	type test struct {
		input  int
		expect int
	}

	tests := []test{
		{
			input:  234,
			expect: 15,
		},
		{
			input:  4421,
			expect: 21,
		},
	}

	fmt.Println("------------------------LeetCode Problem 1281------------------------")
	for _, test := range tests {
		fmt.Printf("Input: %v Output: %v\n", test.input, subtractProductAndSum(test.input))
	}
}
