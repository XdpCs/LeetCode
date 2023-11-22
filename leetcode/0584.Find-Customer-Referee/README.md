# [0584.寻找用户推荐人](https://leetcode.cn/problems/find-customer-referee/)

## 题目描述

表: `Customer`

```text
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
| referee_id  | int     |
+-------------+---------+
在 SQL 中，id 是该表的主键列。
该表的每一行表示一个客户的 id、姓名以及推荐他们的客户的 id。
```

找出那些 没有被 `id = 2` 的客户 **推荐** 的客户的姓名。

以 **任意顺序** 返回结果表。

结果格式如下所示。

**示例 1：**

```text
输入： 
Customer 表:
+----+------+------------+
| id | name | referee_id |
+----+------+------------+
| 1  | Will | null       |
| 2  | Jane | null       |
| 3  | Alex | 2          |
| 4  | Bill | null       |
| 5  | Zack | 1          |
| 6  | Mark | 2          |
+----+------+------------+
输出：
+------+
| name |
+------+
| Will |
| Jane |
| Bill |
| Zack |
+------+
```

## 题目大意

找出那些 没有被 `id = 2` 的客户 **推荐** 的客户的姓名。

## 解题思路

这道题唯一需要注意的一点是`referee_id != 2`
查询条件是用于检索不等于2的数据，但它无法检测NULL值的数据。这是因为在SQL中，`NULL值`与`其他任何值`
进行比较都不会返回`TRUE`或`FALSE`，而是返回`UNKNOWN`，所以我们需要使用`or`进行`referee_id is null`的查询

## 代码

```sql
select name
from Customer
where referee_id != 2
   or referee_id is null;
```

