# [1486.数组异或操作](https://leetcode.cn/problems/xor-operation-in-an-array/)

## 题目描述

给你两个整数，`n` 和 `start` 。

数组 `nums` 定义为：`nums[i] = start + 2*i`（下标从 0 开始）且 `n == nums.length` 。

请返回 `nums` 中所有元素按位异或（XOR）后得到的结果。

**示例 1：**

> 输入：n = 5, start = 0
> 输出：8
> 解释：数组 nums 为 [0, 2, 4, 6, 8]，其中 (0 ^ 2 ^ 4 ^ 6 ^ 8) = 8 。"^" 为按位异或 XOR 运算符。


**示例 2：**

> 输入：n = 4, start = 3
> 输出：8
> 解释：数组 nums 为 [3, 5, 7, 9]，其中 (3 ^ 5 ^ 7 ^ 9) = 8.


**示例 3：**

> 输入：n = 1, start = 7
> 输出：7


**示例 4：**

> 输入：n = 10, start = 5
> 输出：2

**限制**

* `1 <= n <= 1000`
* `0 <= start <= 1000`
* `n == nums.length`

## 题目大意

给你两个整数`n` 和 `start`，数组 `nums` 定义为：`nums[i] = start + 2*i`（下标从 0 开始）且 `n == nums.length` 。
请返回 `nums` 中所有元素按位异或（XOR）后得到的结果。

## 解题思路

### 方法一（模拟）

时间复杂度：O(n)

空间复杂度：O(1)

根据题意，通过`for`循环，计算数组中所有元素按位异或（XOR）后得到的结果。

#### 代码

```go
package leetcode

func xorOperation(n int, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans ^= start + i*2
	}
	return ans
}
```

### 方法二（数学）

时间复杂度：O(1)

空间复杂度：O(1)

根据题意，分析得出题目的公式`start⊕(start+2)⊕(start+4)⊕...⊕(start+2*(n-1))`

根据异或的性质，`∀i∈Z，有 4i⊕(4i+1)⊕(4i+2)⊕(4i+3)=0`
，我们希望往这里凑，所以我们提出公共因子2，题目的公式变为`s⊕(s+1)⊕(s+2)⊕...⊕(s+n-1)+e`，s为 $\lfloor start/2
\rfloor$，由于提出一个会导致，最后1位的二进制丢失，所以e表示的是最后1位的二进制数

我们在分类讨论，结果总共有四种

$$sumXor(x)=
\begin{cases}
x& \text{x=4k,k∈Z}\\
(x-1)⊕x& \text{x=4k+1,k∈Z}\\
(x-2)⊕(x-1)⊕x& \text{x=4k+2,k∈Z}\\
(x-3)⊕(x-2)⊕(x-1)⊕x& \text{x=4k+3,k∈Z}\\
\end{cases}$$

简化一下公式：

$$sumXor(x)=
\begin{cases}
x& \text{x=4k,k∈Z}\\
1& \text{x=4k+1,k∈Z}\\
x+1& \text{x=4k+2,k∈Z}\\
0& \text{x=4k+3,k∈Z}\\
\end{cases}$$

题目需要求的是start开始到n的sumXor，根据异或性质，x⊕y⊕y=x（自反性），
`（sumXor(s-1)⊕sumXor(s+n-1)）*2 + e = 题目所要的结果`

e 是 通过假设 start为偶数，那偶数+偶数=偶数，偶数进行异或，最后1位始终为0，假设start为奇数，奇数+偶数=奇数，偶数个奇数异或为0，奇数个奇数异或为1

#### 代码

```go
package leetcode

func xorOperation(n int, start int) int {
	s := start >> 1
	e := start & n & 1
	return (sumXor(s-1)^sumXor(s+n-1))<<1 | e
}

func sumXor(x int) int {
	if x%4 == 0 {
		return x
	} else if x%4 == 1 {
		return 1
	} else if x%4 == 2 {
		return x + 1
	}
	return 0
}

```