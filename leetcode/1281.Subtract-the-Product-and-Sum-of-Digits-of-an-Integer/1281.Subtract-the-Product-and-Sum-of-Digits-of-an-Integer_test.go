package leetcode

// @Title        1281.Subtract-the-Product-and-Sum-of-Digits-of-an-Integer_test.go
// @Description  1281.Subtract-the-Product-and-Sum-of-Digits-of-an-Integer test
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
			Arg:      234,
			Expected: 15,
		},
		{
			Arg:      4421,
			Expected: 21,
		},
	}

	fmt.Println("------------------------LeetCode Problem 1281------------------------")
	for _, testCase := range testCases {
		result := subtractProductAndSum(testCase.Arg.(int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
