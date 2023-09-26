package leetcode

import "strings"

// @Title        0709.To-Lower-Case.go
// @Description  0709.To-Lower-Case solution
// @Create       XdpCs 2023-09-26 12:23
// @Update       XdpCs 2023-09-26 12:23

func toLowerCase(s string) string {
	result := &strings.Builder{}
	result.Grow(len(s))
	for _, c := range s {
		if c >= 65 && c <= 90 {
			c = c | 32
		}
		result.WriteRune(c)
	}
	return result.String()
}
