package leetcode

// @Title        2586.Count-The-Number-Of-Vowel-Strings-In-Range_test.go
// @Description  2586.Count-The-Number-Of-Vowel-Strings-In-Range test
// @Create       XdpCs 2023-11-23 10:49
// @Update       XdpCs 2023-11-23 10:49

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestVowelStrings(t *testing.T) {
	type arg struct {
		words []string
		left  int
		right int
	}
	testCases := []test.Case{
		{
			Arg: arg{
				words: []string{"are", "amy", "u"},
				left:  0,
				right: 2,
			},
			Expected: 2,
		},
		{
			Arg: arg{
				words: []string{"hey", "aeo", "mu", "ooo", "artro"},
				left:  1,
				right: 4,
			},
			Expected: 3,
		},
	}
	fmt.Println("------------------------LeetCode Problem 2586------------------------")
	for _, tc := range testCases {
		caseArg := tc.Arg.(arg)
		actual := vowelStrings(caseArg.words, caseArg.left, caseArg.right)
		assert.Equal(t, tc.Expected, actual, tc.Print(actual))
	}
}
