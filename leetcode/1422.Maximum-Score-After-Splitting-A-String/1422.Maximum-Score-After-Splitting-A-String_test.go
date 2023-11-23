package leetcode

// @Title        1422.Maximum-Score-After-Splitting-A-String_test.go
// @Description	 1422.Maximum-Score-After-Splitting-A-String test
// @Create       XdpCs 2023-11-23 10:30
// @Update       XdpCs 2023-11-23 10:30

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestMaxScore(t *testing.T) {
	testCases := []test.Case{
		{
			Arg:      "011101",
			Expected: 5,
		},
		{
			Arg:      "00111",
			Expected: 5,
		},
		{
			Arg:      "1111",
			Expected: 3,
		},
	}
	for _, testCase := range testCases {
		actual := maxScore(testCase.Arg.(string))
		assert.Equal(t, testCase.Expected, actual)
	}
}
