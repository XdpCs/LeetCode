package leetcode

// @Title        0035.Search-Insert-Position_test.go
// @Description
// @Create       XdpCs 2023-12-03 19:06
// @Update       XdpCs 2023-12-03 19:06

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestSearchInsertPosition(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	testCases := []test.Case{
		{
			Arg: &args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			Expected: 2,
		},
		{
			Arg: &args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			Expected: 1,
		},
		{
			Arg: &args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			Expected: 4,
		},
	}

	fmt.Println("------------------------LeetCode Problem 0035------------------------")
	for _, testCase := range testCases {
		result := searchInsert(testCase.Arg.(*args).nums, testCase.Arg.(*args).target)
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
