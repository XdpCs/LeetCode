package leetcode

// @Title        0258.Add-Digits_test.go
// @Description  0258.Add-Digits test
// @Create       XdpCs 2023-11-06 12:23
// @Update       XdpCs 2023-11-08 16:47

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestAddDigits(t *testing.T) {
	testCases := []test.Case{
		{
			Arg:      38,
			Expected: 2,
		},
		{
			Arg:      0,
			Expected: 0,
		},
	}

	fmt.Println("------------------------LeetCode Problem 0258------------------------")
	for _, testCase := range testCases {
		result := addDigits(testCase.Arg.(int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
