package test

import (
	"fmt"
	"reflect"
)

// @Title        arg.go
// @Description
// @Create       XdpCs 2023-11-08 16:47
// @Update       XdpCs 2023-11-08 16:47

// Case	Arg's fields must be exported
type Case struct {
	Arg    any
	Expect any
}

func (c *Case) Print(output any) string {
	inputValue := reflect.ValueOf(c.Arg)
	outputValue := reflect.ValueOf(output)
	expectValue := reflect.ValueOf(c.Expect)
	return fmt.Sprintf("Input: %v Output: %+v Expect: %+v\n", printValueWithRecursion(inputValue),
		printValueWithRecursion(outputValue), printValueWithRecursion(expectValue))
}

func printValueWithRecursion(v reflect.Value) string {
	var result string
	switch v.Kind() {
	case reflect.Ptr:
		if !v.IsNil() {
			result += v.Type().Name() + "{"
			result += printValueWithRecursion(v.Elem())
			result += "}"
		}
	default:
		result = fmt.Sprintf("%+v", v)
	}
	return result
}
