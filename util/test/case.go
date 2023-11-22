package test

// @Title        case.go
// @Description  print test case
// @Create       XdpCs 2023-11-08 16:47
// @Update       XdpCs 2023-11-13 10:26

import (
	"fmt"

	print "github.com/XdpCs/print-value"
)

// Case	Arg's fields must be exported
type Case struct {
	Arg      any
	Expected any
}

func (c *Case) Print(output any) string {
	return fmt.Sprintf("Input: %+v Output: %+v Expected: %+v\n", print.Print(c.Arg),
		print.Print(output), print.Print(c.Expected))
}
