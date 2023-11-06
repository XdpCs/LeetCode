package leetcode

// @Title        0709.To-Lower-Case_test.go
// @Description  0709.To-Lower-Case_test solution test
// @Create       XdpCs 2023-09-26 12:23
// @Update       XdpCs 2023-09-26 12:23

import (
	"fmt"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	type test struct {
		input  string
		expect string
	}

	tests := []test{
		{
			input:  "Hello",
			expect: "hello",
		},
		{
			input:  "here",
			expect: "here",
		},
		{
			input:  "LOVELY",
			expect: "lovely",
		},
		{
			input:  "al&phaBET",
			expect: "al&phabet",
		},
	}

	fmt.Println("------------------------LeetCode Problem 0709------------------------")
	for _, test := range tests {
		fmt.Printf("Input: %s Output: %s\n", test.input, toLowerCase(test.input))
	}
}
