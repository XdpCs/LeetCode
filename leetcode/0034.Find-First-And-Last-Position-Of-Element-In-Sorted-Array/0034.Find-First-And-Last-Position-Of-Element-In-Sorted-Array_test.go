package leetcode

// @Title        0034.Find-First-And-Last-Position-Of-Element-In-Sorted-Array_test.go
// @Description
// @Create       XdpCs 2023-12-03 19:20
// @Update       XdpCs 2023-12-03 19:20

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestSearchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	testCases := []test.Case{
		{
			Arg: &args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			Expected: []int{3, 4},
		},
		{
			Arg: &args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			Expected: []int{-1, -1},
		},
		{
			Arg: &args{
				nums:   []int{},
				target: 0,
			},
			Expected: []int{-1, -1},
		},
	}
	fmt.Println("------------------------LeetCode Problem 0034------------------------")
	for _, testCase := range testCases {
		result := searchRange(testCase.Arg.(*args).nums, testCase.Arg.(*args).target)
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
