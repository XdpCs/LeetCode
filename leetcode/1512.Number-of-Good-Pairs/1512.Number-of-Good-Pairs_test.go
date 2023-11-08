package leetcode

// @Title        1512.Number-of-Good-Pairs_test.go
// @Description  1512.Number-of-Good-Pairs_test solution test
// @Create       XdpCs 2023-09-25 16:39
// @Update       XdpCs 2023-11-08 16:47

import (
	"fmt"
	"github.com/XdpCs/leetcode/util/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumIdenticalPairs(t *testing.T) {
	type arg struct {
		Nums []int
	}
	testCases := []test.Case{
		{
			Arg: arg{
				Nums: []int{1, 2, 3, 1, 1, 3},
			},
			Expect: 4,
		},
		{
			Arg: arg{
				Nums: []int{1, 1, 1, 1},
			},
			Expect: 6,
		},
		{
			Arg: arg{
				Nums: []int{1, 2, 3},
			},
			Expect: 0,
		},
	}

	fmt.Println("------------------------LeetCode Problem 1512------------------------")
	for _, testCase := range testCases {
		result := numIdenticalPairs(testCase.Arg.(arg).Nums)
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}
