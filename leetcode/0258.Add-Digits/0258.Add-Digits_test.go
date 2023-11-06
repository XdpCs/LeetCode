package leetcode

// @Title        0258.Add-Digits_test.go
// @Description  0258.Add-Digits test
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
			input:  38,
			expect: 2,
		},
		{
			input:  0,
			expect: 0,
		},
	}

	fmt.Println("------------------------LeetCode Problem 0258------------------------")
	for _, test := range tests {
		fmt.Printf("Input: %v Output: %v\n", test.input, addDigits(test.input))
	}
}
