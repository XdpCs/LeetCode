package leetcode

// @Title        0001.Two-Sum_test.go
// @Description
// @Create       XdpCs 2024-05-05 17:05
// @Update       XdpCs 2024-05-05 17:05

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

type args struct {
	nums   []int
	target int
}

func TestTwoSum(t *testing.T) {
	testCases := []test.Case{
		{
			Arg: &args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			Expected: []int{0, 1},
		},
		{
			Arg: &args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			Expected: []int{1, 2},
		},
		{
			Arg: &args{
				nums:   []int{3, 3},
				target: 6,
			},
			Expected: []int{0, 1},
		},
	}
	fmt.Println("------------------------LeetCode Problem 0001------------------------")
	for _, testCase := range testCases {
		result := twoSum(testCase.Arg.(*args).nums, testCase.Arg.(*args).target)
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
func TestTwoSumHash(t *testing.T) {
	testCases := []test.Case{
		{
			Arg: &args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			Expected: []int{1, 0},
		},
		{
			Arg: &args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			Expected: []int{2, 1},
		},
		{
			Arg: &args{
				nums:   []int{3, 3},
				target: 6,
			},
			Expected: []int{1, 0},
		},
	}
	fmt.Println("------------------------LeetCode Problem 0001------------------------")
	for _, testCase := range testCases {
		result := twoSumHash(testCase.Arg.(*args).nums, testCase.Arg.(*args).target)
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
