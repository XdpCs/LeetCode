# [2236.判断根结点是否等于子结点之和](https://leetcode.cn/problems/root-equals-sum-of-children/)

## 题目描述

给你一个`二叉树`的根结点`root`，该二叉树由恰好`3`个结点组成：根结点、左子结点和右子结点。

如果根结点值等于两个子结点值之和，返回`true`，否则返回`false`。

**示例 1：**

![example1](./images/example1.png)

```markdown
输入：root = [10,4,6]
输出：true
解释：根结点、左子结点和右子结点的值分别是 10 、4 和 6 。
由于 10 等于 4 + 6 ，因此返回 true 。
```

**示例 2：**

![example2](./images/example2.png)

```markdown
输入：root = [5,3,1]
输出：false
解释：根结点、左子结点和右子结点的值分别是 5 、3 和 1 。
由于 5 不等于 3 + 1 ，因此返回 false 。
```

**限制**

* 树只包含根结点、左子结点和右子结点
* `-100 <= Node.val <= 100`

## 题目大意

给你一个根结点，判断根结点的值是否等于左右两个子结点值之和

## 解题思路

时间复杂度：O(1)

空间复杂度：O(1)

判断`root.Val`是否等于`root.Left.Val + root.Right.Val`，如果等于返回`true`，否则返回`false`。

## 代码

```go
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return root.Val == (root.Left.Val + root.Right.Val)
}
```

