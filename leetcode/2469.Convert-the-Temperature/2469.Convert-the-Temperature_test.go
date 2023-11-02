package leetcode

// @Title        2469.Convert-the-Temperature_test.go
// @Description  2469.Convert-the-Temperature solution test
// @Create       XdpCs 2023-09-18 21:05
// @Update       XdpCs 2023-09-18 21:05

import (
	"fmt"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	testCases := []struct {
		celsius float64
	}{
		{36.50},
		{122.11},
	}

	fmt.Println("------------------------LeetCode Problem 2469------------------------")
	for _, testCase := range testCases {
		fmt.Printf("Input: %f Output: %v\n", testCase.celsius, convertTemperature(testCase.celsius))
	}
}
