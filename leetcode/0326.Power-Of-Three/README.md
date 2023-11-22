# [0326.3的幂](https://leetcode.cn/problems/power-of-three/)

## 题目描述

给定一个整数，写一个函数来判断它是否是`3`的幂次方。如果是，返回`true`；否则，返回`false`。

整数`n`是`3`的幂次方需满足：存在整数`x`使得 n == $3^x$

**示例 1：**

> 输入：n = 27
> 输出：true

**示例 2：**

> 输入：n = 0
> 输出：false

**示例 3：**

> 输入：n = 9
> 输出：true

**示例 4：**

> 输入：n = 45
> 输出：false

**限制**

* $-2^{31}$ <= n <= $2^{31}$ - 1

## 题目大意

判断这个数是否是`3`的幂次方。如果是，返回`true`；否则，返回`false`。

## 解题思路

### 方法一（打表）

时间复杂度：O(1)

空间复杂度：O(1)

通过将输入数据的范围限定在`int`类型的范围内，我们可以将所有的`3`的幂次方计算出来，然后直接判断`n`是否在这些数字中。

由于输入的是一个`int`类型，因此需要判断`n`是否大于`0`。

### 代码

```go
package leetcode

import "math"

func isPowerOfThree(n int) bool {
	hashSet := map[int]struct{}{}
	for i := 1; i <= math.MaxInt32; i *= 3 {
		hashSet[i] = struct{}{}
	}
	_, ok := hashSet[n]
	return ok
}
```

### 方法二（数学）

时间复杂度：O(1)

空间复杂度：O(1)

通过将输入数据的范围限定在`int`类型的范围内，我们可以将`3`的幂次方最大值计算出来，然后直接判断`n`是否是这个数字的约数。

由于输入的是一个`int`类型，因此需要判断`n`是否大于`0`。

我们可能开始不知道这个数是多大，可以先用`GetMaxInt32PowerOfThree()`计算出来，然后再判断。

### 代码

```go
package leetcode

import "math"

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

func GetMaxInt32PowerOfThree() int {
	i := 1
	for i <= math.MaxInt32 {
		i *= 3
	}
	return i / 3
}

```
