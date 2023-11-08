package leetcode

// @Title        2469.Convert-the-Temperature_test.go
// @Description  2469.Convert-the-Temperature solution test
// @Create       XdpCs 2023-09-18 21:05
// @Update       XdpCs 2023-11-08 16:47

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XdpCs/leetcode/util/test"
)

func TestConvertTemperature(t *testing.T) {
	type arg struct {
		celsius float64
	}
	testCases := []test.Case{
		{
			Arg: arg{
				celsius: 36.50,
			},
			Expect: []float64{309.65, 97.70},
		},
		{
			Arg: arg{
				celsius: 122.11,
			},
			Expect: []float64{395.26, 251.798},
		},
	}

	fmt.Println("------------------------LeetCode Problem 2469------------------------")
	for _, testCase := range testCases {
		result := convertTemperature(testCase.Arg.(arg).celsius)
		assert.Equal(t, result, testCase.Expect, testCase.Print(result))
	}
}
