# [0231.2的幂](https://leetcode.cn/problems/power-of-two/)

## 题目描述

给你一个整数`n`，请你判断该整数是否是`2`的幂次方。如果是，返回 true ；否则，返回 false 。

如果存在一个整数`x`使得 n == $2^x$ ，则认为`n`是`2`的幂次方。

### 示例 1

> 输入：n = 1
> 输出：true
> 解释： $2^0$ = 1

### 示例 2

> 输入：n = 16
> 输出：true
> 解释： $2^4$ = 16

### 示例 3

> 输入：n = 3
> 输出：false

### 示例 4

> 输入：n = 4
> 输出：true

### 示例 5

> 输入：n = 5
> 输出：false

### 限制

* $-2^{31}$ <= n <= $2^{31}$ - 1

## 题目大意

判断这个数是否是`2`的幂次方。如果是，返回`true`；否则，返回`false`。

## 解题思路

### 方法一（数学）

时间复杂度：O(1)

空间复杂度：O(1)

我们可以直接判断`n`是否是`2`的幂次方。如果是，则`n`的二进制表示中只有最高位是`1`，其余都是`0`，因此判断`n & (n - 1)`
是否等于`0`即可。

由于输入的是一个`int`类型，因此需要判断`n`是否大于`0`。

### 方法一（数学）代码

```go
package leetcode

func isPowerOfTwo(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}
```

### 方法二（打表）

时间复杂度：O(1)

空间复杂度：O(1)

通过将输入数据的范围限定在`int`类型的范围内，我们可以将所有的`2`的幂次方计算出来，然后直接判断`n`是否在这些数字中。

由于输入的是一个`int`类型，因此需要判断`n`是否大于`0`。

### 方法二（打表）代码

```go
package leetcode

import "math"

func isPowerOfTwo(n int) bool {
	hashSet := map[int]struct{}{}
	for i := 1; i <= math.MaxInt32; i *= 2 {
		hashSet[i] = struct{}{}
	}
	_, ok := hashSet[n]
	return ok
}
```

### 方法三（位运算）

时间复杂度：O(1)

空间复杂度：O(1)

通过将输入数据的范围限定在`int`类型的范围内，我们可以将`2`的幂次方最大值计算出来，然后直接判断`n`是否是这个数字的约数。

由于输入的是一个`int`类型，因此需要判断`n`是否大于`0`。

### 代码

```go
package leetcode

func isPowerOfTwo(n int) bool {
	maxPowerOfTwo := 1 << 30
	return (n > 0) && (maxPowerOfTwo%n == 0)
}
```
