package leetcode

// @Title        2469.Convert-the-Temperature_test.go
// @Description  2469.Convert-the-Temperature solution test
// @Create       XdpCs 2023-09-18 21:05
// @Update       XdpCs 2023-11-13 10:20

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestConvertTemperature(t *testing.T) {
	testCases := []test.Case{
		{
			Arg:    36.50,
			Expect: []float64{309.65, 97.70},
		},
		{
			Arg:    122.11,
			Expect: []float64{395.26, 251.798},
		},
	}

	fmt.Println("------------------------LeetCode Problem 2469------------------------")
	for _, testCase := range testCases {
		result := convertTemperature(testCase.Arg.(float64))
		assert.Equal(t, testCase.Expect, result, testCase.Print(result))
	}
}
