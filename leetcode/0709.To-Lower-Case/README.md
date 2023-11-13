# [0709. 转换成小写字母](https://leetcode.cn/problems/to-lower-case/)

## 题目描述

给你一个字符串 `s` ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。

**示例 1：**

> 输入：s = "Hello"
> 输出："hello"

**示例 2：**

> 输入：s = "here"
> 输出："here"

**示例 3：**

> 输入：s = "LOVELY"
> 输出："lovely"

**限制**

* `1 <= s.length <= 100`
* `s` 由**ASCII**字符集中的可打印字符组成

## 题目大意

给你一个字符串 `s` ，将该字符串中的大写字母转换成相同的小写字母，不是大写字母，按照原样输出，返回新的字符串

## 解题思路

时间复杂度：O(n)

空间复杂度：O(n)

这个题目，第一感觉是用 `strings.ToLower()` 函数，但是这样就没有意思了。

很多人看到这个题，第一个想法是使用 `strings.Builder`
类型来拼接字符串，然后遍历字符串，如果是`大写字母`就加`32`转换成`小写字母`，然后拼接到 `strings.Builder` 类型中

其实可以使用位运算来实现，大写字母和小写字母的 ASCII 码相差 `32`，`32`的二进制为`0100000`，`65`的二进制为`1000001`，`90`
的二进制为`1011010`，根据观察二进制，我们可以使用`或运算`实现加法，这里用到了 `strings.Builder`
来拼接字符串，这个类型的 `WriteRune()` 方法可以直接写入 rune 类型的字符，不需要转换。

## 代码

```go
package leetcode

import "strings"

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
```
