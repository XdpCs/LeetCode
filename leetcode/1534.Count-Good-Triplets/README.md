# [1534. 统计好三元组](https://leetcode.cn/problems/count-good-triplets/)

## 题目描述

给你一个整数数组 `arr` ，以及 `a`、`b` 、`c` 三个整数。请你统计其中好三元组的数量。

如果三元组 `(arr[i], arr[j], arr[k])` 满足下列全部条件，则认为它是一个**好三元组**。

* `0 <= i < j < k < arr.length`
* `|arr[i] - arr[j]| <= a`
* `|arr[j] - arr[k]| <= b`
* `|arr[i] - arr[k]| <= c`

其中 `|x|` 表示 `x` 的绝对值。

返回 **好三元组的数量**。

**示例 1：**

```text
输入：arr = [3,0,1,1,9,7], a = 7, b = 2, c = 3
输出：4
解释：一共有 4 个好三元组：[(3,0,1), (3,0,1), (3,1,1), (0,1,1)] 。
```

**示例 2：**

```text
输入：arr = [1,1,2,2,3], a = 0, b = 0, c = 1
输出：0
解释：不存在满足所有条件的三元组。
```

**限制**

* `3 <= arr.length <= 100`
* `0 <= arr[i] <= 1000`
* `0 <= a, b, c <= 1000`

## 题目大意

给你一个整数数组 `arr` ，以及 `a`、`b` 、`c` 三个整数。如果三元组 `(arr[i], arr[j], arr[k])` 满足下列全部条件，则认为它是一个
**好三元组**。

* `0 <= i < j < k < arr.length`
* `|arr[i] - arr[j]| <= a`
* `|arr[j] - arr[k]| <= b`
* `|arr[i] - arr[k]| <= c`

统计三元组的数量。

## 解题思路

时间复杂度：O($n^3$)

空间复杂度：O(1)

根据题意，通过使用三层`for`循环，遍历所有可能的三元组，判断是否满足条件，如果满足条件则将结果加一，可以通过在第二层`for`
中判断，进行剪枝，减少不必要的循环。

## 代码

```go
package leetcode

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if abs(arr[i]-arr[j]) <= a {
				for k := j + 1; k < len(arr); k++ {
					if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
						count++
					}
				}
			}
		}
	}
	return count
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
```
