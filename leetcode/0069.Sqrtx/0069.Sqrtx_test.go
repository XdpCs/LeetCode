package leetcode

// @Title        0069.Sqrtx_test.go
// @Description
// @Create       XdpCs 2023-12-04 16:59
// @Update       XdpCs 2023-12-04 16:59

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestSqrtX(t *testing.T) {
	testCases := []test.Case{
		{
			Arg:      4,
			Expected: 2,
		},
		{
			Arg:      8,
			Expected: 2,
		},
	}
	fmt.Println("------------------------LeetCode Problem 0069------------------------")
	for _, testCase := range testCases {
		result := mySqrt(testCase.Arg.(int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
