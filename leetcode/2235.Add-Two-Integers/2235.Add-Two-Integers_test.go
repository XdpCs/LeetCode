package leetcode

// @Title        2235.Add-Two-Integers_test.go
// @Description  2235.Add-Two-Integers solution test
// @Create       XdpCs 2023-09-18 21:05
// @Update       XdpCs 2023-11-08 16:47

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestSum(t *testing.T) {
	type arg struct {
		Num1 int
		Num2 int
	}
	testCases := []test.Case{
		{
			Arg: arg{
				Num1: 1118,
				Num2: 1114,
			},
			Expect: 1118 + 1114,
		},
		{
			Arg: arg{
				Num1: 12,
				Num2: 5,
			},
			Expect: 12 + 5,
		},
		{
			Arg: arg{
				Num1: -10,
				Num2: 4,
			},
			Expect: -10 + 4,
		},
		{
			Arg: arg{
				Num1: 0,
				Num2: 0,
			},
			Expect: 0,
		},
	}

	fmt.Println("------------------------LeetCode Problem 2235------------------------")
	for _, testCase := range testCases {
		arg := testCase.Arg.(arg)
		result := sum(arg.Num1, arg.Num2)
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}
