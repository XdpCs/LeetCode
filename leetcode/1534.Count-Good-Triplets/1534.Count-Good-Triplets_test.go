package leetcode

// @Title        1534.Count-Good-Triplets_test.go
// @Description  1534.Count-Good-Triplets_test solution test
// @Create       XdpCs 2023-09-25 19:12
// @Update       XdpCs 2023-11-08 16:47

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestCountGoodTriplets(t *testing.T) {
	type arg struct {
		Arr []int
		A   int
		B   int
		C   int
	}
	testCases := []test.Case{
		{
			Arg: arg{
				Arr: []int{3, 0, 1, 1, 9, 7},
				A:   7,
				B:   2,
				C:   3,
			},
			Expect: 4,
		},
		{
			Arg: arg{
				Arr: []int{1, 1, 2, 2, 3},
				A:   0,
				B:   0,
				C:   1,
			},
			Expect: 0,
		},
	}

	fmt.Println("------------------------LeetCode Problem 1534------------------------")
	for _, testCase := range testCases {
		arg := testCase.Arg.(arg)
		result := countGoodTriplets(arg.Arr, arg.A, arg.B, arg.C)
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}
