# [1757.可回收且低脂的产品](https://leetcode.cn/problems/recyclable-and-low-fat-products/)

## 题目描述

表：`Products`

```text
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| product_id  | int     |
| low_fats    | enum    |
| recyclable  | enum    |
+-------------+---------+
product_id 是该表的主键（具有唯一值的列）。
low_fats 是枚举类型，取值为以下两种 ('Y', 'N')，其中 'Y' 表示该产品是低脂产品，'N' 表示不是低脂产品。
recyclable 是枚举类型，取值为以下两种 ('Y', 'N')，其中 'Y' 表示该产品可回收，而 'N' 表示不可回收。
```

编写解决方案找出既是低脂又是可回收的产品编号。

返回结果**无顺序要求**。

返回结果格式如下例所示：

### 示例 1

```text
输入：
Products 表：
+-------------+----------+------------+
| product_id  | low_fats | recyclable |
+-------------+----------+------------+
| 0           | Y        | N          |
| 1           | Y        | Y          |
| 2           | N        | Y          |
| 3           | Y        | Y          |
| 4           | N        | N          |
+-------------+----------+------------+
输出：
+-------------+
| product_id  |
+-------------+
| 1           |
| 3           |
+-------------+
解释：
只有产品 id 为 1 和 3 的产品，既是低脂又是可回收的产品。
```

## 题目大意

找出既是低脂又是可回收的产品编号

## 解题思路

这道题需要知道低脂是`low_fats = 'Y'`，可回收是`recyclable = 'Y'`，即可以得出结果

## 代码

```sql
select product_id
from Products
where low_fats = 'Y'
  and recyclable = 'Y';
```
