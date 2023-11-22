package leetcode

// @Title        0263.Ugly-Number_test.go
// @Description  0263.Ugly-Number test
// @Create       XdpCs 2023-11-22 20:25
// @Update       XdpCs 2023-11-22 20:25

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestIsUgly(t *testing.T) {
	testCases := []test.Case{
		{
			Arg:      6,
			Expected: true,
		},
		{
			Arg:      1,
			Expected: true,
		},
		{
			Arg:      14,
			Expected: false,
		},
	}

	fmt.Println("------------------------LeetCode Problem 0263------------------------")
	for _, testCase := range testCases {
		result := isUgly(testCase.Arg.(int))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
