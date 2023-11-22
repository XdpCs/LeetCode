package leetcode

// @Title        0709.To-Lower-Case_test.go
// @Description  0709.To-Lower-Case_test solution test
// @Create       XdpCs 2023-09-26 12:23
// @Update       XdpCs 2023-11-08 16:47

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestToLowerCase(t *testing.T) {

	testCases := []test.Case{
		{
			Arg:      "Hello",
			Expected: "hello",
		},
		{
			Arg:      "here",
			Expected: "here",
		},
		{
			Arg:      "LOVELY",
			Expected: "lovely",
		},
		{
			Arg:      "al&phaBET",
			Expected: "al&phabet",
		},
	}

	fmt.Println("------------------------LeetCode Problem 0709------------------------")
	for _, testCase := range testCases {
		result := toLowerCase(testCase.Arg.(string))
		assert.Equal(t, testCase.Expected, result, testCase.Print(result))
	}
}
