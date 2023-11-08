package leetcode

// @Title        1486.XOR-Operation-in-an-Array_test.go
// @Description  1486.XOR-Operation-in-an-Array_test solution test
// @Create       XdpCs 2023-09-25 16:17
// @Update       XdpCs 2023-11-08 16:47

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

type arg struct {
	N     int
	Start int
}

var testCases = []test.Case{
	{
		Arg: arg{
			N:     5,
			Start: 0,
		},
		Expect: 8,
	},
	{
		Arg: arg{
			N:     4,
			Start: 3,
		},
		Expect: 8,
	},
	{
		Arg: arg{
			N:     1,
			Start: 7,
		},
		Expect: 7,
	},
	{
		Arg: arg{
			N:     10,
			Start: 5,
		},
		Expect: 2,
	},
}

func TestXorOperation(t *testing.T) {
	fmt.Println("------------------------LeetCode Problem 1486------------------------")
	for _, testCase := range testCases {
		arg := testCase.Arg.(arg)
		result := xorOperation(arg.N, arg.Start)
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}

func TestXorOperationMath(t *testing.T) {
	fmt.Println("------------------------LeetCode Problem 1486------------------------")
	for _, testCase := range testCases {
		arg := testCase.Arg.(arg)
		result := xorOperationMath(arg.N, arg.Start)
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}
