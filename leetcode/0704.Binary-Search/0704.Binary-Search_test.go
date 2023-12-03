package leetcode

// @Title        0704.Binary-Search_test.go
// @Description
// @Create       XdpCs 2023-12-03 18:01
// @Update       XdpCs 2023-12-03 18:01

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	testCases := []test.Case{
		{
			Arg: &args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			Expected: 4,
		},
		{
			Arg: &args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			Expected: -1,
		},
	}

	fmt.Println("------------------------LeetCode Problem 0704------------------------")
	for _, testCase := range testCases {
		result := search(testCase.Arg.(*args).nums, testCase.Arg.(*args).target)
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
